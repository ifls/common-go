package orm

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

type Account struct {
	Id      int64  `xorm:"pk autoincr 'aid'"`
	Name    string `xorm:"unique"`
	Balance float64
	Version string
}

type Book struct {
	Id      int64 `xorm:"bid"`
	Name    string
	Price   float64
	Version int `xorm:"version"`
}

var testCli SqlCli

func init() {
	//host := "47.107.151.251"
	host := "23.91.101.147"
	password := "#Wfs123456"
	driver := "mysql"
	url := MakeMysqlUrl("root", password, host, 3306, "test")
	cli, err := Open(driver, url)
	if err != nil {
		log.Fatal(err)
	}
	testCli = cli
}

func TestConnStatus(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := testCli.PingContext(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlCli_CreateTable(t *testing.T) {
	acc1 := Account{
		Id:      1,
		Name:    "1",
		Balance: 1,
		Version: "1.1",
	}

	err := testCli.CreateTable(acc1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlCli_Query(t *testing.T) {
	rows, err := testCli.Query("SELECT * FROM test.account LIMIT 0,3")
	if err != nil {
		t.Fatal(err)
	}

	for i, kvs := range rows {
		log.Printf("index =%d, kvs =%s\n", i, kvs)
	}
}

func TestSqlCli_Exec(t *testing.T) {
	err := testCli.Exec(fmt.Sprintf("update test.account set balance = %d where aid = %d", 14, 2))
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlCli_Insert(t *testing.T) {

	acc1 := Account{
		Name:    strconv.Itoa(rand.Int() % 10000),
		Balance: rand.Float64() * 100,
		Version: "1.1",
	}
	acc2 := Account{
		Name:    strconv.Itoa(rand.Int() % 10000),
		Balance: rand.Float64() * 100,
		Version: "1.2",
	}
	err := testCli.Insert(acc1, acc2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlCli_GetOneRow(t *testing.T) {
	acc1 := Account{
		Id: 1,
	}

	err := testCli.Get(&acc1)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%#v\n", acc1)
}

func TestSqlCli_Exist(t *testing.T) {
	acc1 := Account{
		Id: 5,
	}

	err := testCli.Exist(&acc1)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%#v\n", acc1)
}

func TestSqlCli_Find(t *testing.T) {
	accs := make([]Account, 0)

	err := testCli.Find(&accs)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%#v\n", accs)
}

func TestSqlCli_Iterate(t *testing.T) {
	handd := func(idx int, bean interface{}) error {
		log.Printf("index = %d, bean = %#v\n", idx, bean)
		return nil
	}
	acc1 := Account{}

	err := testCli.Iterate(acc1, handd)
	if err != nil {
		t.Fatal(err)
	}

	log.Printf("%#v\n", acc1)
}

func TestSqlCli_Update(t *testing.T) {
	acc1 := Account{
		Name: "updated_name",
	}
	err := testCli.Update(acc1, 3)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlCli_Count(t *testing.T) {
	acc1 := Account{}
	count, err := testCli.Count(acc1)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf(" count = %d\n", count)
}

func TestSqlCli_Sum(t *testing.T) {
	acc1 := Account{}
	sum, err := testCli.Sum(acc1, "balance")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("sum = %v\n", sum)
}

func TestSqlCli_Delete(t *testing.T) {
	acc1 := Account{}
	err := testCli.Delete(acc1, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSqlCli_DeleteTable(t *testing.T) {
	acc1 := Account{}
	err := testCli.DeleteTable(acc1)
	if err != nil {
		t.Fatal(err)
	}
}
