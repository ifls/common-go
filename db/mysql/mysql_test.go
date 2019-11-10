package mysql

import (
	"log"
	"testing"
	//pb "server_go/struct/proto"
)

type MysqlUser2 struct {
	Id   int
	Name string `orm:"size(100)"`
}

type MysqlUser struct {
	Uid      int64  `db:"uid"`
	ID       string `db:"ID"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	Password string `db:"password"`
}

func init() {
	host := ""
	port := 99999
	_ = MakeMysqlUrl("root", "02#F20ebac", host, uint16(port), "test")
	log.Println(new(MysqlUser))
	log.Println(new(MysqlUser2))
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
