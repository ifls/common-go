package mongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

//const mongoUrlFormat = "mongodb://%s:%d"

type Client struct {
	*mongo.Client
}

var timeout int = 5

func init() {

}

//返回连接指定url的mongo客户端
func NewClient(mongoUrl string) (Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return Client{}, fmt.Errorf("mongo.NewClient err =%w, url = %s\n", err, mongoUrl)
	}

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		return Client{}, fmt.Errorf("client.connect err =%w, url = %s\n", err, mongoUrl)
	}

	//err = client.Ping(ctx, readpref.Primary())
	//if err != nil {
	//	return Client{}, fmt.Errorf("client.Ping err =%w, url = %s\n", err, mongoUrl)
	//}
	return Client{client}, nil
}

//支持插入，一行或者多行， 指定数据库和表名会自动创建
func (c Client) Insert(dbName string, tableName string, docs ...interface{}) error {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	cols := make([]interface{}, 0)
	cols = append(cols, docs...)

	if len(docs) > 1 {
		ret, err := col.InsertMany(ctx, cols)
		if err != nil {
			return fmt.Errorf("insertOne err = %w", err)
		}

		if ret == nil {
			return fmt.Errorf("insert id is error")
		}

		if len(ret.InsertedIDs) != len(docs) {
			return fmt.Errorf("want insert %d, but insert %d, error", len(docs), len(ret.InsertedIDs))
		}

	} else if len(docs) == 1 {
		ret, err := col.InsertOne(ctx, docs[0])
		if err != nil {
			return fmt.Errorf("insertOne err = %w\n", err)
		}
		if ret == nil {
			return fmt.Errorf("insert id is error")
		}
	} else {
		return errors.New("insert params error")
	}

	return nil
}

//返回多行迭代器
func (c Client) FindMany(dbName string, tableName string, filter interface{}) (*mongo.Cursor, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	cur, err := col.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("findMany err =%w\n", err)
	}
	return cur, nil
}

//返回单行结果
func (c Client) FindOne(dbName string, tableName string, filter interface{}) (*mongo.SingleResult, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	//一个都找不到也是error
	ret := col.FindOne(ctx, filter)
	if ret.Err() != nil {
		return nil, fmt.Errorf("findOne err = %w\n", ret.Err())
	}
	return ret, nil
}

func (c Client) Update(dbName string, tableName string, filter interface{}, update interface{}, onlyOne bool) (*mongo.UpdateResult, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	var ret *mongo.UpdateResult
	var err error
	if onlyOne {
		ret, err = col.UpdateMany(ctx, filter, update)
	} else {
		ret, err = col.UpdateMany(ctx, filter, update)
	}

	if err != nil {
		return nil, fmt.Errorf("update err = %w", err)
	}

	return ret, nil
}

func (c Client) ReplaceOne(dbName string, tableName string, filter interface{}, replace interface{}) (*mongo.UpdateResult, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	ret, err := col.ReplaceOne(ctx, filter, replace)
	if err != nil {
		return nil, fmt.Errorf("raplace err = %w", err)
	}

	return ret, nil
}

//删除所有匹配的
func (c Client) Delete(dbName string, tableName string, filter interface{}, onlyOne bool) (int64, error) {
	col := c.Database(dbName).Collection(tableName)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)

	var ret *mongo.DeleteResult
	var err error
	if onlyOne {
		ret, err = col.DeleteMany(ctx, filter)
	} else {
		ret, err = col.DeleteMany(ctx, filter)
	}

	if err != nil {
		return 0, fmt.Errorf("delete err = %w", err)
	}

	return ret.DeletedCount, nil
}
