package redis

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/assert"
	"log"
	"strconv"
	"testing"
)

func init() {
	ip := "23.91.101.147"
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
	defer func() {
		_ = c.Close()
	}()

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
	//hmGet(c, "hash_name", "name1", "code1")
}

func TestKGet(t *testing.T) {
	key := "test_key"
	val := "3333"
	err := KSet(key, val)
	if err != nil {
		t.Fatalf("redis set error, err=%v", err)
	}

	val2, err := KGet(key)
	if err != nil {
		t.Fatalf("redis get error, err=%v", err)
	}
	//util.DevInfo("%s:[%s]\n", key, val)
	assert.Equal(t, "3333", string(val2.([]uint8)), "get value is diff")
}

func TestKSet(t *testing.T) {
	key := "test_key"
	val := "3333"
	err := KSet(key, val)
	if err != nil {
		t.Fatalf("redis set error, err=%v", err)
	}

	val2, err := KGet(key)
	if err != nil {
		t.Fatalf("redis get error, err=%v", err)
	}

	val2Str := string(val2.([]uint8))
	assert.Equal(t, val, val2Str, "redis set value err, get[%v] != set[%v]", val2, val)
}

func TestKDel(t *testing.T) {
	key := "test_del_key"
	val := "3333"
	err := KSet(key, val)
	if err != nil {
		t.Fatalf("redis set error, err=%v", err)
	}

	val2, err := KGet(key)
	if err != nil {
		t.Fatalf("redis get error, err=%v", err)
	}
	//log.Printf("%s\n", val2)
	val2Str := string(val2.([]uint8))
	assert.Equal(t, val, val2Str, "redis set value err, get[%v] != set[%v]", val2, val)

	err = KDel(key)
	if err != nil {
		t.Fatal(err)
	}

	val3, err := KGet(key)
	if err != nil {
		t.Fatal(err)
	}
	//log.Printf("%s\n", val3)
	assert.Equal(t, nil, val3)
}

func TestHMSet(t *testing.T) {
	kvs := make([]interface{}, 0)
	kvs = append(kvs, "name1")
	kvs = append(kvs, "he2")
	kvs = append(kvs, "code1")
	kvs = append(kvs, "yifeng2")
	err := HMSet("test_hash", kvs...)
	if err != nil {
		t.Fatal(err)
	}

	keys := make([]interface{}, 0)
	keys = append(keys, "name1")
	keys = append(keys, "code1")

	reply, err := HMGet("test_hash", keys...)
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range reply {
		//util.DevInfo("%d:%s\n", i, v)
		//util.DevInfo("%s\n", kvs[2*i + 1])
		v2 := string(v.([]byte))
		assert.Equal(t, kvs[2*i+1], v2, "hmset error value is diff")
	}
}

func TestHMGet(t *testing.T) {
	reply, err := HMGet("test_hash", "name1", "code1")
	if err != nil {
		t.Fatal(err)
	}
	vals := []string{"he2", "yifeng2"}
	//vals array
	for i, v := range reply {
		//util.DevInfo("%v:%s", i, v)
		assert.Equal(t, vals[i], string(v.([]uint8)), "hmget value is diff")
	}
}

func TestHGetAll(t *testing.T) {
	reply, err := HGetAll("test_all_hash")
	if err != nil {
		t.Fatal(err)
	}

	vals := []string{"key1", "val1", "key2", "val2"}
	//vals array
	for i, v := range reply {
		//util.DevInfo("%v:%s", i, v)
		assert.Equal(t, vals[i], string(v.([]uint8)), "hmget value is diff")
	}
}

func TestIncr(t *testing.T) {
	key := "newbee"
	err := KSet(key, 0)
	if err != nil {
		t.Fatal(err)
	}
	reply, err := Incr(key)
	if err != nil {
		t.Fatal(err)
	}

	val, err := KGet(key)
	if err != nil {
		t.Fatal(err)
	}
	intval, err := strconv.Atoi(string(val.([]uint8)))
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, int64(intval), reply)
}

func TestTemp(t *testing.T) {
	err := HMSet("test_all_hash", "key1", "val1", "key2", "val2")
	if err != nil {
		t.Fatal(err)
	}
}
