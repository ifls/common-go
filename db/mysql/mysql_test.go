package mysql

import (
	"testing"
	//pb "server_go/struct/proto"
)

type MysqlUser2 struct {
	Id   int
	Name string `orm:"size(100)"`
}

const (
	CREATE_TABLE_USER = "create table proto(uid int unique key auto_increment, password varchar(20), ID varchar(18));"
)

type MysqlUser struct {
	uid      int64  `db:"uid"`
	ID       string `db:"ID"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	password string `db:"password"`
}

func init() {
	//host := ""
	//port := 99999
	//url := MakeMysqlUrl("root", "02#F20ebac", host, uint16(port), "test")
}

func TestMysqlConnection(t *testing.T) {
	//err := testClient.Ping()
	//if err != nil {
	//	t.Fatalf("mysql database cannot connect")
	//}
}

func TestInsert(t *testing.T) {
	//err := insert(defaultClient)
	//if err != nil {
	//	t.Fatalf("mysql database cannot insert or insert fail")
	//}
}

func TestStats(t *testing.T) {
	//stat := testClient.Stats()
	//util.DevInfo("%v", stat)

}
