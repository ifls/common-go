package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
	"gocore/util"
	"log"
	"testing"
)

var ip string = "47.107.151.251"

func init() {
	redisUrl = ip + ":6379"
	_, err := Open(redisUrl)
	if err != nil {
		log.Println(err)
		return
	}
}

func TestRedisConn(t *testing.T) {
	c, err := redis.Dial("tcp", redisUrl)
	if err != nil {
		fmt.Printf("redis conn error = %v", err)
		return
	}
	//推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用
	//多个defer，按照后进先出
	defer c.Close()

	////c.Send("get", )
	//val, err := kGet(c, "aa")
	//if val == "" {
	//	fmt.Printf("get kk nil \n")
	//} else {
	//	fmt.Printf("get kk=%s \n", val)
	//}
	//
	//
	//kSet(c, "aa", strconv.Itoa(111))
	//
	//val, err = kGet(c, "aa")
	//if val == "" {
	//	fmt.Printf("get kk nil \n")
	//} else {
	//	fmt.Printf("get kk=%s \n", val)
	//}

	//err = hmset(c, "hash_name", "name1", "he", "code1", "yifeng")
	hmget(c, "hash_name", "name1", "code1")
}

func TestKGet(t *testing.T) {
	val, err := KGet("test_key")
	if err != nil {
		util.LogErr(err, zap.String("reason", "redis get"))
		t.Fatalf("redis get error, err=%v", err)
	}
	util.DevInfo("[%s]\n", val)
}

func TestKSet(t *testing.T) {
	vv := "3333"
	err := KSet("test_key", vv)
	if err != nil {
		util.LogErr(err, zap.String("reason", "redis get"))
		t.Fatalf("redis set error, err=%v", err)
	}

	val, err := KGet("test_key")
	if err != nil {
		util.LogErr(err, zap.String("reason", "redis get"))
		t.Fatalf("redis get error, err=%v", err)
	}

	val2 := string(val.([]uint8))
	if val2 != vv {
		t.Fatalf("redis set value err, get[%v] != set[%v]", val2, vv)
	}
}

func TestHMGet(t *testing.T) {
	reply, err := HMGet("test_hash", "name1", "code1")
	util.LogErrAndExit(err)

	for i, v := range reply {
		util.DevInfo("%d:%d", i, v)
	}
}

func TestHMSet(t *testing.T) {
	keys := make([]interface{}, 0)
	keys = append(keys, "name2")
	keys = append(keys, "code2")

	kvs := make([]interface{}, 0)
	kvs = append(kvs, "name2")
	kvs = append(kvs, "he2")
	kvs = append(kvs, "code2")
	kvs = append(kvs, "yifeng2")
	err := HMSet("test_hash", kvs...)
	util.LogErrAndExit(err)

	reply, err := HMGet("test_hash", keys...)
	util.LogErrAndExit(err)

	for i, v := range reply {
		util.DevInfo("%d:%s\n", i, v)
		util.DevInfo("%s\n", kvs[2*i+1])
		v2 := string(v.([]byte))
		if v2 != kvs[2*i+1] {
			t.Fatal("value is diff")
		}
	}
}

func TestIncr(t *testing.T) {
	KSet("newbee", 0)
	Incr("newbee")
	//reply, err := Incr("newbee")
	//t.Log("")
}
