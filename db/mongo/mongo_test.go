package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"gocore/util"
	"testing"
	"time"
)

type MongoUser struct {
	Uid       int
	Name      string
	Password  string
	LoginTime string
}

type MongoBook struct {
	Bid  int
	Name string
}

func TestMongoInsertOne(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.InsertOne(ctx, MongoUser{
		Uid:       int(util.NextId()) % 1000000,
		Name:      "user100332",
		Password:  "www",
		LoginTime: time.Now().Format(util.TIME_FORMAT),
	})
	if err != nil {
		util.LogErr(err)
	}

	util.DevInfo("inserted id = %v", ret.InsertedID)
}

func TestCreateCollection(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("book")
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	ret, err := collection.InsertOne(ctx, MongoBook{
		Bid:  int(util.NextId()) % 1000000,
		Name: "book100332",
	})
	if err != nil {
		util.LogErr(err)
	}

	util.DevInfo("inserted id = %v", ret.InsertedID)
}

func TestMongoApiInsertOne(t *testing.T) {
	user := MongoUser{
		Uid:       int(util.NextId()) % 1000000,
		Name:      "user100332",
		Password:  "www",
		LoginTime: time.Now().Format(util.TIME_FORMAT),
	}

	err := insertMongo(testMongoClient, "test", "user", user)
	if err != nil {
		util.LogErr(err)
	}
}

func TestMongoApiInsert(t *testing.T) {
	users := []interface{}{
		MongoUser{
			Uid:       int(util.NextId()) % 1000000,
			Name:      "user100332",
			Password:  "www",
			LoginTime: time.Now().Format(util.TIME_FORMAT),
		},
		MongoUser{
			Uid:       int(util.NextId()) % 1000000,
			Name:      "user100332",
			Password:  "www",
			LoginTime: time.Now().Format(util.TIME_FORMAT),
		},
		MongoUser{
			Uid:       int(util.NextId()) % 1000000,
			Name:      "user100332",
			Password:  "www",
			LoginTime: time.Now().Format(util.TIME_FORMAT),
		},
	}

	err := insertMongo(testMongoClient, "test", "user", users...)
	if err != nil {
		util.LogErr(err)
	}
}

func TestMongoInsert(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	users := []interface{}{
		MongoUser{
			Uid:       int(util.NextId()) % 1000000,
			Name:      "user100332",
			Password:  "www",
			LoginTime: time.Now().Format(util.TIME_FORMAT),
		},
		MongoUser{
			Uid:       int(util.NextId()) % 1000000,
			Name:      "user100332",
			Password:  "www",
			LoginTime: time.Now().Format(util.TIME_FORMAT),
		},
		MongoUser{
			Uid:       int(util.NextId()) % 1000000,
			Name:      "user100332",
			Password:  "www",
			LoginTime: time.Now().Format(util.TIME_FORMAT),
		},
	}

	ret, err := collection.InsertMany(ctx, users)
	if err != nil {
		util.LogErr(err)
	}

	util.DevInfo("insert result = %+v", ret)
}

func TestMongoFindOne(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ret2 := collection.FindOne(ctx, bson.M{"name": "user100332"})
	if ret2.Err() != nil {
		util.LogErr(ret2.Err())
	}
	var user1 MongoUser
	if err := ret2.Decode(&user1); err != nil {
		util.LogErr(err)
	}
	util.DevInfo("%+v\n", user1)
}

func TestMongoApiFindOne(t *testing.T) {
	ret, err := FindOneMongo(testMongoClient, "test", "user", bson.M{"name": "user100332"})
	if err != nil {
		util.LogErr(err)
	}
	var user1 MongoUser
	if err := ret.Decode(&user1); err != nil {
		util.LogErr(err)
	}
	util.DevInfo("%+v\n", user1)
}

func TestMongoApiFind(t *testing.T) {
	cur, err := FindManyMongo(testMongoClient, "test", "user", bson.M{"name": "user100332"})
	if err != nil {
		util.LogErr(err)
		return
	}

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	for cur.Next(ctx) {
		var user1 MongoUser
		if err := cur.Decode(&user1); err != nil {
			util.LogErr(err)
		}
		util.DevInfo("%+v\n", user1)
	}
}

func TestMongoFind(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cur, err := collection.Find(ctx, bson.M{"name": "user100332"})
	if err != nil {
		util.LogErr(err)
		return
	}

	for cur.Next(ctx) {
		var user1 MongoUser
		if err := cur.Decode(&user1); err != nil {
			util.LogErr(err)
		}
		util.DevInfo("%+v\n", user1)
	}
}

func TestMongoUpdateOne(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.UpdateOne(ctx, bson.M{"name": "user100332"}, bson.M{"$set": bson.M{"name": "UserName_changed"}})
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("update one result %+v\n", ret)
}

func TestMongoUpdate(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.UpdateMany(ctx, bson.M{"name": "user100332"}, bson.M{"$set": bson.M{"name": "UserName_changed"}})
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("update one result %+v\n", ret)
}

func TestMongoApiUpdate(t *testing.T) {
	err := UpdateMongo(testMongoClient, "test", "user", bson.M{"name": "user100332"}, bson.M{"$set": bson.M{"name": "UserName_changed"}})
	if err != nil {
		util.LogErr(err)
		return
	}
}

func TestMongoReplace(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	user := MongoUser{
		Uid:       int(util.NextId()) % 1000000,
		Name:      "user100332",
		Password:  "www",
		LoginTime: time.Now().Format(util.TIME_FORMAT),
	}

	ret, err := collection.ReplaceOne(ctx, bson.M{"name": "UserName_changed"}, user)
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("update one result %+v\n", ret)
}

func TestMongoDeleteOne(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.DeleteOne(ctx, bson.M{"name": "user100332"})
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("update one result %+v\n", ret)
}

func TestMongoDelete(t *testing.T) {
	collection := testMongoClient.Database("test").Collection("user")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := collection.DeleteMany(ctx, bson.M{"name": "UserName_changed"})
	if err != nil {
		util.LogErr(err)
		return
	}

	util.DevInfo("update one result %+v\n", ret)
}

func TestMongoApiDelete(t *testing.T) {
	err := DeleteMongo(testMongoClient, "test", "user", bson.M{"name": "user100332"})
	if err != nil {
		util.LogErr(err)
		return
	}
}
