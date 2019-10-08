package consul

import (
	"errors"
	"github.com/hashicorp/consul/api"
	"log"
)

func Get(key string) ([]byte, error) {
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	//log.Printf("KV: %#v\n queryMeta = %#v", pair, qm)

	if pair == nil {
		return nil, errors.New("key not exists")
	}

	return pair.Value, nil
}

//递归遍历所有目录, 没有时即是空数组
func List(keyprefix string) (api.KVPairs, error) {
	pairs, qm, err := kv.List(keyprefix, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("pairs = %#v, qm = %#v\n", pairs, qm)
	if pairs == nil {
		return make(api.KVPairs, 0), nil
	}

	for i, pair := range pairs {
		log.Printf("i = %d, KV : %v %s\n", i, pair.Key, pair.Value)
	}
	return pairs, nil
}

//递归遍历所有目录, 相比List, 只获得key数组
func Keys(keyPrefix string) ([]string, error) {
	keys, qm, err := kv.Keys(keyPrefix, "", nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	log.Printf("keys = %#v, qm = %v", keys, qm)
	if keys == nil {
		return make([]string, 0), nil
	}
	return keys, nil
}

//幂等
func Put(key string, val []byte) error {
	pair := &api.KVPair{Key: key, Value: val}
	_, err := kv.Put(pair, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	//log.Printf("KV PUT: %v %s\n writeMeta = %#v", pair.Key, pair.Value, writeMeta)
	return nil
}

//幂等, 只能以/结尾的目录可以且只可以删除空目录
func Delete(key string) error {
	wm, err := kv.Delete(key, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("writeMeta = %#v", wm)
	return nil
}

//删除整个目录， 不必须以/结尾
func DeleteTree(prefix string) error {
	wm, err := kv.DeleteTree(prefix, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("writeMeta=%#v\n", wm)
	return nil
}

//必须拿之前get过的kvpair，来指定cas， 如果中间没有其他修改才会执行， 否则返回失败，需要不断重试
func DeleteCAS(pair *api.KVPair) error {
	succ, writeMeta, err := kv.DeleteCAS(pair, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("succ, pair = %#v %#v\n writeMeta = %#v", succ, pair, writeMeta)
	if succ != true {
		return errors.New("delete cas fail ")
	}
	return nil
}

func Txn(txns api.KVTxnOps) error {
	succ, txnRsp, writeMeta, err := kv.Txn(txns, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("succ = %#v, txnRsp=%#v, wm = %#v", succ, txnRsp, writeMeta)
	return nil
}

//get到的pair, 如果没有被修改过就可以拿新的kpair覆盖
func CAS(pair *api.KVPair) error {
	succ, wm, err := kv.CAS(pair, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("succ = %v, wm = %v", succ, wm)
	if succ != true {
		return errors.New("CAS fail")
	}
	return nil
}

func Acquire(pair *api.KVPair) error {
	succ, wm, err := kv.Acquire(pair, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("succ =%#v, writeMeta=%#v\n", succ, wm)

	if succ != true {
		return errors.New("acquire lock-session fail")
	}

	return nil
}

func Release(pair *api.KVPair) error {
	succ, wm, err := kv.Release(pair, nil)
	if err != nil {
		log.Print(err)
		return err
	}
	log.Printf("succ =%#v, writeMeta=%#v\n", succ, wm)

	if succ != true {
		return errors.New("release lock-session fail")
	}

	return nil
}
