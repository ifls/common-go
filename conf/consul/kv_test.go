package consul

import (
	"github.com/hashicorp/consul/api"
	"gopkg.in/go-playground/assert.v1"
	"log"
	"math/rand"
	"strconv"
	"testing"
)

func TestApiGet(t *testing.T) {
	//get 已存在的key
	val, err := Get("test_get_key")
	if err != nil {
		t.FailNow()
	}
	assert.Equal(t, "test_get_value", string(val))

	//get 不存在的key
	val2, err := Get("test_get_nxt_key")
	assert.NotEqual(t, nil, err)
	if err != nil {
		log.Print(err)
		assert.Equal(t, nil, val2)
		t.FailNow()
	}

}

func TestApiList(t *testing.T) {
	kvs, err := List("")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, kvs)

	kvs, err = List("t")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, kvs)

	kvs, err = List("gggg")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, kvs)
}

func TestApiKeys(t *testing.T) {
	keys, err := Keys("")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, keys)

	keys, err = Keys("t")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, keys)

	keys, err = Keys("gggg")
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEqual(t, nil, keys)
}

func TestApiPut(t *testing.T) {
	key := "test_put_key"
	val := "test_put_value2"
	err := Put(key, []byte(val))
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}

	//get test
	val2, err := Get(key)
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
	assert.Equal(t, val, string(val2))
}

func TestApiDelete(t *testing.T) {
	key := "test_delete_key"
	err := Delete(key)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}

	//get test
	val, err := Get(key)
	assert.NotEqual(t, nil, err)
	if err != nil {
		log.Print(err)
		assert.Equal(t, nil, val)
		//t.FailNow()
	}

	key = "test_delete_dir2/"
	err = Delete(key)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}

func TestAPiDeleteTree(t *testing.T) {
	key := "test_delete_dir"
	err := DeleteTree(key)
	if err != nil {
		log.Print(err)
		t.FailNow()
	}
}

func TestApiDeleteCAS(t *testing.T) {
	pair, _, err := kv.Get("test_delete_key", nil)
	if err != nil {
		t.Fatal(err)
	}
	if pair != nil {
		err = DeleteCAS(pair)
		if err != nil {
			log.Print(err)
			t.FailNow()
		}
	}

	_, err = kv.Put(&api.KVPair{
		Key:   "test_delete_key",
		Value: []byte("test_delete_val"),
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	pair, _, err = kv.Get("test_delete_key", nil)
	if err != nil {
		t.Fatal(err)
	}

	if pair != nil {
		_, err = kv.Put(&api.KVPair{
			Key:   "test_delete_key",
			Value: []byte("test_delete_val2"),
		}, nil)

		if err != nil {
			log.Fatal(err)
		}

		err = DeleteCAS(pair)
		if err != nil {
			log.Print(err)
			return
		}
		t.Fatal(err)
	}
}

func TestApiTxn(t *testing.T) {
	txns := api.KVTxnOps{
		&api.KVTxnOp{
			Verb:  api.KVSet,
			Key:   "test_txn_put_key",
			Value: []byte("test_txn_put_val"),
		},
		&api.KVTxnOp{
			Verb: api.KVDelete,
			Key:  "test_txn_put_key",
		},
	}
	err := Txn(txns)
	if err != nil {
		t.Fatal(err)
	}
}

func TestApiCAS(t *testing.T) {
	key := "test_cas_key"
	pair, qm, err := kv.Get(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n qm = %#v\n", pair, qm)

	//_, err = kv.Put(&api.KVPair{
	//		Key:   key,
	//		Value: []byte("test_delete_val2"),
	//	}, nil)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	if pair != nil {
		pair.Value = []byte("changed4")
		err := CAS(pair)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestApiAcquire(t *testing.T) {
	key := "test_cas_key"
	pair, qm, err := kv.Get(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n qm = %#v\n", pair, qm)

	err = Acquire(pair)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	key := "test_get_key"
	pair, qm, err := kv.Get(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("KV: %v %s\n queryMeta = %#v", pair.Key, pair.Value, qm)
	assert.Equal(t, "test_get_value", string(pair.Value))
}

//[]string keys
func TestKeys(t *testing.T) {
	key := "host/local"
	keys, qm, err := kv.Keys(key, "", nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("keys %s\n queryMeta = %#v", keys, qm)
}

//kv pairs
func TestList(t *testing.T) {
	key := "host/local"
	kvs, qm, err := kv.List(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	for i, v := range kvs {
		log.Printf("index=%v key=%#v value=%s \n", i, v.Key, v.Value)
	}
	log.Printf("qm = %#v\n", qm)
}

func TestSetAndGet(t *testing.T) {
	key := "test_put_key"
	val := strconv.Itoa(rand.Int())
	p := &api.KVPair{Key: key, Value: []byte(val)}
	_, err := kv.Put(p, nil)
	if err != nil {
		t.Fatal(err)
	}

	pair, _, err := kv.Get(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("get KV: %v %s\n", pair.Key, pair.Value)
	assert.Equal(t, val, string(pair.Value))
}

func TestDeleteCAS(t *testing.T) {
	pair, qm, err := kv.Get("test_delete_key", nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n qm = %#v\n", pair, qm)

	succ, wm, err := kv.DeleteCAS(pair, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("succ = %v, wm = %v", succ, wm)
}

func TestCAS(t *testing.T) {
	key := "test_cas_key"
	pair, qm, err := kv.Get(key, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n qm = %#v\n", pair, qm)

	_, err = kv.Put(&api.KVPair{
		Key:   key,
		Value: []byte("test_delete_val2"),
	}, nil)

	if err != nil {
		log.Fatal(err)
	}

	if pair != nil {
		pair.Value = []byte("changed3")
		succ, wm, err := kv.CAS(pair, nil)
		if err != nil {
			t.Fatal(err)
		}
		log.Printf("succ = %v, wm = %v", succ, wm)
	}
}

//锁的是什么？
func TestAcquire(t *testing.T) {
	//key := "test_cas_key"
	//pair, qm, err := kv.Get(key, nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//log.Printf("%#v\n qm = %#v\n", pair, qm)

	id, wm, err := session.CreateNoChecks(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer session.Destroy(id, nil)
	log.Printf("id =%#v, writeMeta=%#v\n", id, wm)

	pair := &api.KVPair{
		Key:     "test_acquire2",
		Value:   []byte("test_acquire_value"),
		Session: id,
	}
	succ, wm, err := kv.Acquire(pair, nil)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("succ =%#v, writeMeta=%#v\n", succ, wm)
	//time.Sleep(200 * time.Second)
}
