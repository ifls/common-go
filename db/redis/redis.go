package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/ifls/gocore/util"
)

var redisClient redis.Conn

var redisUrl string

func Open(addr string) (redis.Conn, error) {
	cli, err := redis.Dial("tcp", addr)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}

	redisUrl = addr
	redisClient = cli
	return cli, nil
}

//GetRedisUrl
//func GetRedisUrl() string {
//	return redisUrl
//}

///// inner api /////////////////
//
//func incrBy() {
//
//}

func kSet(c redis.Conn, key string, value interface{}) error {
	_, err := c.Do("SET", key, value)
	if err != nil {
		return fmt.Errorf("redis kSet() %w", err)
	}

	//util.DevInfo("%v, %s\n", reply, reply)
	return nil
}

func kGet(c redis.Conn, key string) (interface{}, error) {
	reply, err := c.Do("GET", key)
	if err != nil {
		return nil, fmt.Errorf("redis kGet() %w", err)
	}

	//util.DevInfo("%v, %s\n", reply, reply)
	return reply, nil
}

func kDel(c redis.Conn, key string) error {
	_, err := c.Do("DEL", key)
	if err != nil {
		return fmt.Errorf("redis kDel() %w", err)
	}

	//util.DevInfo("%v, %s\n", reply, reply)
	return nil
}

//func hSet() {
//
//}
//
//func hDel() {
//
//}
//
//func hClear() {
//
//}

func hmSet(c redis.Conn, hash string, kvs ...interface{}) error {
	argvs := redis.Args{}.Add(hash).Add(kvs...)

	_, err := c.Do("HMSET", argvs...)
	if err != nil {
		return fmt.Errorf("redis hmSet() %w", err)
	}

	//util.DevInfo("vals = %s\n", reply)
	return nil
}

func hGetAll(c redis.Conn, hash string) ([]interface{}, error) {
	vals, err := redis.Values(c.Do("HGETALL", hash))
	if err != nil {
		return nil, fmt.Errorf("redis hGetAll() %w", err)
	}
	return vals, nil
}

//func hget(c redis.Conn, key string) (interface{}, error) {
//	return nil, nil
//}

func hmGet(c redis.Conn, hash string, keys ...interface{}) ([]interface{}, error) {
	argvs := redis.Args{}.Add(hash).Add(keys...)

	vals, err := redis.Values(c.Do("HMGET", argvs...))
	if err != nil {
		return nil, fmt.Errorf("redis hmGet() %w", err)
	}

	//util.DevInfo("vals =  %s\n", reply)
	return vals, nil
}

func incr(c redis.Conn, key string) (interface{}, error) {
	reply, err := c.Do("INCR", key)
	if err != nil {
		return nil, fmt.Errorf("redis hmGet() %w", err)
	}

	//util.DevInfo("%v, %s\n", reply, reply)
	return reply, nil
}

func commonCheck(cli redis.Conn) error {
	if cli == nil {
		return fmt.Errorf("cli conn == nil")
	}
	return nil
}

////////////////// redis public API///////////////////////

func KSet(key string, value interface{}) error {
	if err := commonCheck(redisClient); err != nil {
		return err
	}
	return kSet(redisClient, key, value)
}

func KGet(key string) (interface{}, error) {
	if err := commonCheck(redisClient); err != nil {
		return nil, err
	}
	return kGet(redisClient, key)
}

func KDel(key string) error {
	if err := commonCheck(redisClient); err != nil {
		return err
	}
	return kDel(redisClient, key)
}

func HMSet(hash string, kvs ...interface{}) error {
	if err := commonCheck(redisClient); err != nil {
		return err
	}
	return hmSet(redisClient, hash, kvs...)
}

func HMGet(hash string, keys ...interface{}) ([]interface{}, error) {
	if err := commonCheck(redisClient); err != nil {
		return nil, err
	}
	return hmGet(redisClient, hash, keys...)
}

// return [k,v,k,v]
func HGetAll(hash string) ([]interface{}, error) {
	if err := commonCheck(redisClient); err != nil {
		return nil, err
	}
	return hGetAll(redisClient, hash)
}

//func HSet(hash string, key string, value interface{}) error {
//	if err := commonCheck(redisClient); err != nil {
//		return err
//	}
//	return nil
//}
//
//func HGet(hash string, key string) interface{} {
//	if err := commonCheck(redisClient); err != nil {
//		return err
//	}
//	return nil
//}

func Incr(key string) (interface{}, error) {
	if err := commonCheck(redisClient); err != nil {
		return nil, err
	}
	return incr(redisClient, key)
}
