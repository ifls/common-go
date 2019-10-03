package tidb

import (
	"fmt"
	"testing"
	//pb "server_go/struct/proto"
)

type TestUser2 struct {
	Id   int
	Name string `orm:"size(100)"`
}

type TestUser struct {
	uid      int64  `db:"uid"`
	ID       string `db:"ID"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	password string `db:"password"`
}

func TestInsertData(t *testing.T) {
	pool = open()
	defaultSetting(pool)
	Ping()
	//insert(pool)
	//query(pool, "ss")
}

func TestConnection(t *testing.T) {
	err := testClient.Ping()
	if err != nil {
		t.Fatalf("tidb database cannot connect")
	}

}

func TestMysqlConnection(t *testing.T) {
	pool = open()
	err := pool.Ping()
	if err != nil {
		t.Fatalf("tidb database cannot connect")
	}

}

func TestInsert(t *testing.T) {
	pool = open()

	err := insert(pool)
	if err != nil {
		t.Fatalf("tidb database cannot insert or insert fail")
	}
}

func TestStats(t *testing.T) {
	pool = open()
	stat := pool.Stats()
	fmt.Printf("%v", stat)

}

func init() {
	// set default database
	//orm.RegisterDataBase("default", "tidb", "username:password@tcp(127.0.0.1:3306)/db_name?charset=utf8", 30)
	//
	//// register model
	//orm.RegisterModel(new(User))
	//
	//// create table
	//orm.RunSyncdb("default", false, true)
	//
	//o := orm.NewOrm()
	//
	//user := User2{Name: "jj"}
	//
	//// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//// update
	//user.Name = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//// read one
	//u := User2{Id: user.Id}
	//err = o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//
	//// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
