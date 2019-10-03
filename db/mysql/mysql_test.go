package mysql

import (
	"gocore/util"
	"testing"
	//pb "server_go/struct/proto"
)

type MysqlUser2 struct {
	Id   int
	Name string `orm:"size(100)"`
}

func TestMysqlConnection(t *testing.T) {
	err := testClient.Ping()
	if err != nil {
		t.Fatalf("mysql database cannot connect")
	}
}

func TestInsert(t *testing.T) {
	err := insert(defaultClient)
	if err != nil {
		t.Fatalf("mysql database cannot insert or insert fail")
	}
}

func TestStats(t *testing.T) {
	stat := testClient.Stats()
	util.DevInfo("%v", stat)

}
