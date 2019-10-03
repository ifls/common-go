package conf

import (
	"context"
	"errors"
	"github.com/coreos/etcd/clientv3"
	"gocore/util"
	"time"
)

var defaultClient *clientv3.Client
var testClient *clientv3.Client
var err error

var host string = "192.168.8.101"

func init() {
	defaultClient, err = connectEtcd()
	testClient = defaultClient
}

func connectEtcd() (*clientv3.Client, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{host + ":2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		util.LogErr(err)
		return nil, err
	}
	return cli, nil
}

func put(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := defaultClient.Put(ctx, key, value)
	cancel()
	if err != nil {
		util.LogErr(err)
		return err
	}

	util.DevInfo("%+v\n", resp)
	return nil
}

func get(key string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := defaultClient.Get(ctx, key)
	cancel()
	if err != nil {
		util.LogErr(err)
		return nil, err
	}

	util.DevInfo("%+v\n", resp.Kvs)
	for _, v := range resp.Kvs {
		util.DevInfo("%s\n", v.Value)
		return v.Value, nil
	}
	return nil, errors.New("key not found value")
}
