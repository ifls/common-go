package mongo

import (
	"context"
	"github.com/ifls/gocore/utils"
	log2 "github.com/ifls/gocore/utils/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strconv"
	"testing"
	"time"
)

type MongoUser struct {
	Uid       int
	Name      string
	Password  string
	LoginTime string
}

var testMongoClient Client
var testDb string = "test"

func init() {
	url := "mongodb://47.107.151.251:27017"
	client, err := NewClient(url)
	if err != nil {
		log2.LogErr(err)
	}
	testMongoClient = client
}

func TestMongoApiInsert(t *testing.T) {
	rd := int(utils.NextId()) % 1000000
	user := MongoUser{
		Uid:       rd,
		Name:      "user" + strconv.Itoa(rd),
		Password:  "www",
		LoginTime: time.Now().Format(utils.TimeFormat),
	}
	user2 := user
	user2.Password = "ccc"
	err := testMongoClient.Insert(testDb, "user", user, user2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMongoApiFindOne(t *testing.T) {
	rd := int(utils.NextId()) % 1000000
	user := MongoUser{
		Uid:       rd,
		Name:      "user" + strconv.Itoa(rd),
		Password:  "www",
		LoginTime: time.Now().Format(utils.TimeFormat),
	}
	err := testMongoClient.Insert(testDb, "user", user)
	if err != nil {
		t.Fatal(err)
	}

	ret, err := testMongoClient.FindOne(testDb, "user", bson.M{"name": "user" + strconv.Itoa(rd)})
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("mongo.singleResult = %#v\n", ret)
	var user1 MongoUser
	if err := ret.Decode(&user1); err != nil {
		log2.LogErr(err)
	}
	log2.DevInfo("%+v\n", user1)
}

func TestMongoApiFind(t *testing.T) {
	cur, err := testMongoClient.FindMany(testDb, "user", bson.M{"password": "www"})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("Cursor = %#v\n", cur)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	for cur.Next(ctx) {
		var user1 MongoUser
		if err := cur.Decode(&user1); err != nil {
			log2.LogErr(err)
		}
		log2.DevInfo("%+v\n", user1)
	}
}

func TestMongoUpdateOne(t *testing.T) {
	collection := testMongoClient.Database(testDb).Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.UpdateOne(ctx, bson.M{"name": "user100332"}, bson.M{"$set": bson.M{"name": "UserName_changed"}})
	if err != nil {
		t.Fatal(err)
	}

	log2.DevInfo("update one result %+v\n", ret)
}

func TestMongoUpdateMany(t *testing.T) {
	collection := testMongoClient.Database(testDb).Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.UpdateMany(ctx, bson.M{"name": "user100332"}, bson.M{"$set": bson.M{"name": "UserName_changed"}})
	if err != nil {
		t.Fatal(err)
	}

	log2.DevInfo("update one result %+v\n", ret)
}

func TestMongoApiUpdate(t *testing.T) {
	ret, err := testMongoClient.Update(testDb, "user", bson.M{"name": "user100332"}, bson.M{"$set": bson.M{"name": "UserName_changed"}}, false)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("updateResult = %#v\n", ret)
}

func TestMongoReplaceOne(t *testing.T) {
	user := MongoUser{
		Uid:       int(utils.NextId()) % 1000000,
		Name:      "user100332x",
		Password:  "www",
		LoginTime: time.Now().Format(utils.TimeFormat),
	}

	ret, err := testMongoClient.ReplaceOne(testDb, "user", bson.M{"password": "www"}, user)
	if err != nil {
		t.Fatal(err)
	}

	log2.DevInfo("replaceOne %+v\n", ret)
}

func TestMongoDeleteOne(t *testing.T) {
	collection := testMongoClient.Database(testDb).Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.DeleteOne(ctx, bson.M{"name": "user100332"})
	if err != nil {
		t.Fatal(err)
	}

	log2.DevInfo("update one result %+v\n", ret)
}

func TestMongoDelete(t *testing.T) {
	collection := testMongoClient.Database(testDb).Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.DeleteMany(ctx, bson.M{"name": "UserName_changed"})
	if err != nil {
		t.Fatal(err)
	}

	log2.DevInfo("update one result %+v\n", ret)
}

func TestMongoApiDelete(t *testing.T) {
	ret, err := testMongoClient.Delete(testDb, "user", bson.M{"name": "user100332"}, false)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("deleteResult =  %#v\n", ret)
}

func TestMongoAggregate(t *testing.T) {
	collection := testMongoClient.Database(testDb).Collection("log")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	curs, err := collection.Aggregate(ctx, mongo.Pipeline{bson.D{
		{"$group", bson.D{
			{"_id", nil},
			{"max", bson.D{
				{"$max", "$logid"},
			}},
		}},
	}})
	if err != nil {
		t.Fatal(err)
	}

	log2.DevInfo("update one result %+v\n", curs)
	if curs == nil {
		t.Fatal("nil")
	}
	//var data struct{}
	//err = curs.All(ctx, &data)
	//if err != nil {
	//	t.Fatalf("all %s\n", err)
	//}
	//log.Printf("data %+v\n", data)
	for curs.Next(ctx) {
		var user1 struct {
			_id string
			Max int
		}

		if err := curs.Decode(&user1); err != nil {
			log2.LogErr(err)
		}
		log2.DevInfo("%+v\n", user1)
	}
}
