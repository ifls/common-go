package mongo

import (
	"context"
	"errors"
	"fmt"
	"github.com/ifls/gocore/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
	"time"
)

//const mongoUrlFormat = "mongodb://%s:%d"

var mongoUrl string

var testMongoClient *mongo.Client

//var defaultMongoClient *mongo.Client

func init() {
	client := prepare()
	testMongoClient = client
	//defaultMongoClient = client
}

//连接指定环境mongo客户端
func prepare() *mongo.Client {
	mongoUrl = ""
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		util.LogErr(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		util.LogErr(err, zap.String("reason", "ping error mongo url = "+mongoUrl))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		util.LogErr(err, zap.String("reason", "ping error mongo url = "+mongoUrl))
	}
	return client
}

//支持插入，一行和多行， 指定数据库和表名会自动创建
func insertMongo(c *mongo.Client, dbName string, tableName string, docs ...interface{}) error {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cols := make([]interface{}, 0)
	cols = append(cols, docs...)

	if len(docs) > 1 {
		ret, err := col.InsertMany(ctx, cols)
		if err != nil {
			util.LogErrReason(err, "insertMany error")
			return err
		}
		util.LogInfo(fmt.Sprintf("%+v\n", ret))
	} else if len(docs) == 1 {
		ret, err := col.InsertOne(ctx, docs[0])
		if err != nil {
			util.LogErrReason(err, "insertOne error")
			return err
		}
		util.LogInfo(fmt.Sprintf("%+v\n", ret))
	} else {
		return errors.New("insert params error")
	}

	return nil
}

//返回多行迭代器
func FindManyMongo(c *mongo.Client, dbName string, tableName string, filter interface{}) (*mongo.Cursor, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := col.Find(ctx, filter)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}
	return cur, nil
}

//返回单行结果
func FindOneMongo(c *mongo.Client, dbName string, tableName string, filter interface{}) (*mongo.SingleResult, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)

	ret := col.FindOne(ctx, filter)
	if ret.Err() != nil {
		util.LogErrReason(ret.Err(), "FindOne error")
		return nil, ret.Err()
	}
	return ret, nil
}

func UpdateMongo(c *mongo.Client, dbName string, tableName string, filter interface{}, update interface{}) error {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := col.UpdateMany(ctx, filter, update)
	if err != nil {
		util.LogErr(err)
		return err
	}
	util.DevInfo("update one result %+v\n", ret)
	return nil
}

func DeleteMongo(c *mongo.Client, dbName string, tableName string, filter interface{}) error {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	ret, err := col.DeleteMany(ctx, filter)
	if err != nil {
		util.LogErr(err)
		return err
	}

	util.DevInfo("update one result %+v\n", ret)
	return nil
}

//func dispose() {
//
//}
