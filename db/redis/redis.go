package redis

import (
	"github.com/gomodule/redigo/redis"
	"gocore/util"
)

var redisClient redis.Conn

var redisUrl string

func init() {
	redisUrl = ""
	redisClient, _ = Open(redisUrl)
}

func Open(addr string) (redis.Conn, error) {
	cli, err := redis.Dial("tcp", addr)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}

	return cli, nil
}

func incrBy() {

}

func kSet(c redis.Conn, key string, value interface{}) error {
	reply, err := c.Do("SET", key, value)
	if err != nil {
		util.LogErr(err)
		return err
	}

	util.DevInfo("%s", reply)
	return nil
}

func kGet(c redis.Conn, key string) (interface{}, error) {
	reply, err := c.Do("GET", key)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}

	util.DevInfo("%s\n", reply)
	return reply, nil
}

func kDel() {

}

func hSet() {

}

func hDel() {

}

func hClear() {

}

func hmset(c redis.Conn, hash string, kvs ...interface{}) error {
	argvs := redis.Args{}.Add(hash).Add(kvs...)

	reply, err := c.Do("HMSET", argvs...)
	if err != nil {
		util.LogErr(err)
		return err
	}

	util.DevInfo("%v, %s\n", reply, reply)
	return nil
}

func hget(c redis.Conn, key string) (interface{}, error) {
	return nil, nil
}

func hmget(c redis.Conn, hash string, keys ...interface{}) ([]interface{}, error) {
	argvs := redis.Args{}.Add(hash).Add(keys...)

	reply, err := c.Do("HMGET", argvs...)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}

	util.DevInfo("%v, %s\n", reply, reply)
	return reply.([]interface{}), nil
}

func incr(c redis.Conn, key string) (interface{}, error) {
	reply, err := c.Do("INCR", key)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}

	util.DevInfo("%v, %s\n", reply, reply)
	return reply, nil
}

////////////////// redis API///////////////////////

func KSet(key string, value interface{}) error {
	return kSet(redisClient, key, value)
}

func KGet(key string) (interface{}, error) {
	return kGet(redisClient, key)
}

func HMSet(hash string, kvs ...interface{}) error {
	return hmset(redisClient, hash, kvs...)
}

func Hset(hash string, key string, value interface{}) error {
	return nil
}

func HGet(hash string, key string) interface{} {
	return nil
}

func HMGet(hash string, keys ...interface{}) ([]interface{}, error) {
	return hmget(redisClient, hash, keys...)
}

func Incr(key string) (interface{}, error) {
	return incr(redisClient, key)
}
