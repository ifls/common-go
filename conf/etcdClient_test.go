package conf

import (
	"github.com/coreos/etcd/clientv3"
	"gocore/util"
	"gopkg.in/go-playground/assert.v1"
	"testing"
	"time"
)

//var host string = "192.168.8.101"

func TestEtcdConnect(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host + ":2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		util.LogErr(err)
		return
	}
	defer cli.Close()
}

func TestEtcdPut(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host + ":2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		util.LogErr(err)
		return
	}
	defer cli.Close()

	key := "testKey"
	val := "testVal"
	put(key, val)

	data, _ := get(key)

	assert.Equal(t, string(data), val)
}

func TestEtcdGet(t *testing.T) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host + ":2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		util.LogErr(err)
		return
	}
	defer cli.Close()

	key := "/rpcx_log/LogRpcServer"

	data, _ := get(key)

	util.DevInfo("key = %s, val = [%s]", key, string(data))
}
