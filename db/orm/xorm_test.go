package orm

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/ifls/gocore/util"
	"go.uber.org/zap"
	"testing"
	"time"
)

type Account struct {
	Id      int64
	Name    string `xorm:"unique"`
	Balance float64
	Version int `xorm:"version"`
}

type Book struct {
	Id      int64
	Name    string
	Price   float64
	Version int
}

var host string
var dburl string
var x *xorm.Engine

func init() {
	host = ""
	dburl = ""

	var err error
	x, err = connectDb(dburl)
	if err != nil {
		util.DevInfo("connect error" + err.Error())
	}
}

func TestXORM(t *testing.T) {
	t.Log(host)
	//mysqlUrl := dburl

	orm, err := xorm.NewEngine("mysql", dburl)
	if err != nil {
		util.LogErr(err, zap.String("reason", "new engine"))
		return
	}

	if err = orm.Ping(); err != nil {
		util.LogErr(err, zap.String("reason", "ping "))
		return
	}

	////自动检测和创建表
	//if err := orm.Sync2(new(Account)); err != nil {
	//	util.LogErr(err)
	//}

	if err := orm.Sync2(new(Book)); err != nil {
		util.LogErr(err)
	}

	tables, err := orm.DBMetas()
	if err != nil {
		util.LogErr(err)
		return
	}

	for _, table := range tables {
		t.Logf("%v\n", table.Type)
		t.Logf("tablename=%s\n", table.Name)
	}

	//acc := new(Account)
	//acc.Name = "hh19"
	//acc.Balance = 99
	//acc.Version = 1
	//affacted, err := orm.Insert(acc)
	//if err != nil {
	//	util.LogErr(err)
	//	return
	//}
	//t.Logf("%d\n", affacted)

	book := new(Book)
	book.Name = "linux"
	affacted, err := orm.Insert(book)
	if err != nil {
		util.LogErr(err)
		return
	}
	t.Logf("%d\n", affacted)

	bk := Account{}
	has, err := orm.Id(1).Get(&bk)
	if err != nil {
		util.LogErr(err)
		return
	}

	if has {
		t.Logf("%v\n", bk)
	}

	//user := Account{}
	//has, err := orm.Id(3).Get(&user)
	//if err != nil {
	//	util.LogErr(err)
	//	return
	//}
	//
	//if has {
	//	t.Logf("%v\n", user)
	//}
}

func TestConnect(t *testing.T) {
	if err := x.Ping(); err != nil {
		t.Fatal("connect error" + err.Error())
	}
}

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(j).Format("2006-01-02 15:04:05.000000")), nil
}

type Testtableg struct {
	Id         int64  `xorm:"autoincr pk"`
	Name       string `xorm:"notnull "`
	CreateTime int64  `xorm:"created"`
	UpdateTime int64  `xorm:updated`
}

func TestCreateTable(t *testing.T) {
	//if err := CreateTable(&gostruct.FileCore{}); err != nil {
	//	t.Fatal("create table" + err.Error())
	//}
}

func TestInsert(t *testing.T) {
	//book1 := new(gostruct.FileCore)
	//
	//book1.CreateTime = time.Now().Format(util.TIME_FORMAT)
	//book1.OssUrl = "ss"
	//
	//err := InsertRecord(book1)
	//if err != nil {
	//	t.Fatal("insert error")
	//}
}

func TestExec(t *testing.T) {
	sql := "select * from testtabled"
	results, err := Exec(sql)
	if err != nil {
		util.LogErr(err)
	}
	for k, v := range results {
		util.DevInfo("%v:%s\n", k, v)
	}
}

func TestQueryRecord(t *testing.T) {
	sql := "select * from testtabled"
	results, err := Exec(sql)
	if err != nil {
		util.LogErr(err)
	}
	for k, v := range results {
		util.DevInfo("%v:%s\n", k, v)
	}
}

func TestDeleteTable(t *testing.T) {
	//book1 := new(gostruct.FileCore)
	//err := DeleteTable(book1)
	//if err != nil {
	//	t.Fatal("insert error")
	//}
}

func TestGet(t *testing.T) {
	book1 := new(Testtableg)
	has, err := Get(book1, "Id=?", 1)
	if err != nil {
		util.LogErr(err)
	}
	t.Logf("%v %v\n", has, book1)
}

func TestFind(t *testing.T) {
	book1 := make([]Testtableg, 0)
	err := Find(&book1, "Id>?", 0)
	if err != nil {
		t.Logf("err = %s\n", err.Error())
	}
	for i, row := range book1 {
		t.Logf("%d, %v\n", i, row)
	}
}
