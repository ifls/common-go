package ssdb

import (
	"fmt"
	"github.com/ssdb/gossdb/ssdb"
	"os"
)

var (
	host = "192.168.0.113"
	port = 8888
)

func init() {
}

func ssdbConn() {
	db, err := ssdb.Connect(host, port)
	if err != nil {
		fmt.Printf("ssdb conn error:%v", err)
		os.Exit(1)
	}

	var val interface{}
	_, _ = db.Set("a", "xxx")
	val, err = db.Get("a")
	fmt.Printf("%s\n", val)
	_, _ = db.Del("a")
	val, err = db.Get("a")
	fmt.Printf("%s\n", val)

	_, _ = db.Do("zset", "z", "a", 3)
	_, _ = db.Do("multi_zset", "z", "b", -2, "c", 5, "d", 3)
	resp, err := db.Do("zrange", "z", 0, 10)
	if err != nil {
		os.Exit(1)
	}
	if len(resp)%2 != 1 {
		fmt.Printf("bad response")
		os.Exit(1)
	}

	fmt.Printf("Status: %s\n", resp[0])
	for i := 1; i < len(resp); i += 2 {
		fmt.Printf("  %s : %3s\n", resp[i], resp[i+1])
	}
	return
}
