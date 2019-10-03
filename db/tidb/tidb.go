package tidb

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gocore/util"
	"time"
)

const (
	TABLE_USER = "create table proto(uid int unique key auto_increment, password varchar(20), ID varchar(18));"
)

var pool *sql.DB

type MysqlConn struct {
	user     string
	password string
	net      string
	host     string
	port     uint16
	database string
}

var tidbUrl string
var tidbUrlFormatter string = "%s:%s@%s(%s:%d)/%s"

var testClient *sql.DB
var defaultClient *sql.DB

func init() {
	tidbUrl = ""
	db := open()
	if db != nil {
		testClient = db
		defaultClient = db
	}
}

func defaultSetting(mysql *sql.DB) {
	mysql.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close, 过期的连接在下次使用时才会被lazyclose,
	mysql.SetMaxOpenConns(100)                  //设置最大连接数
	mysql.SetMaxIdleConns(16)                   //设置闲置连接赤的最大值
}

func Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	//defer cancel()

	err := pool.PingContext(ctx)
	if err != nil {
		fmt.Printf("unable to connect to database: %v", err)
	}
	return err
}

func open() (db *sql.DB) {
	db, err := sql.Open("mysql", tidbUrl)
	if err != nil {
		util.LogErr(err)
		return nil
	}
	return db
}

func createConn(db *sql.DB) {

}

func close(db *sql.DB) {

}

func createTable(db *sql.DB) {

}

func alterTable(db *sql.DB) {

}

func tableScheme(db *sql.DB) {

}

func deleteTable(db *sql.DB) {

}

//只执行,不返回
func exec(sql string) {

}

func insert(db *sql.DB) error {
	result, err := db.Exec("insert into proto(password, id) values(?,?)", "password1", "430223199509050711")
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return err
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return err
	}

	if lastInsertID < 1 {
		return fmt.Errorf("%s", "insert id < 1")
	}

	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected() //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return err
	}

	if rowsaffected < 1 {
		//return fmt.Errorf("%s", "insert infect zero row")
		return errors.New("insert affect zero row")
	}

	fmt.Println("RowsAffected:", rowsaffected)
	return nil
}

//返回一个sql语句,以后执行
func _prepare(db *sql.DB) {

}

//queryRow执行一个sql,期望返回至多一行
//执行,并返回结果
func query(db *sql.DB, sql string) {
	//user := new(TestUser)
	//row := db.QueryRow("select * from proto where uid = 3")
	////row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	//if err := row.Scan(&user.uid, &user.ID, &user.password); err != nil {
	//	fmt.Printf("scan failed, err:%v", err)
	//	return
	//}
	//fmt.Printf("row : %v\n", *user)
}

func queryWithTimeOut(db *sql.DB) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//user := new(TestUser)
	//row := db.QueryRowContext(ctx, "select * from proto where uid = 3")
	////row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	//if err := row.Scan(&user.uid, &user.ID, &user.password); err != nil {
	//	fmt.Printf("scan failed, err:%v", err)
	//	return
	//}
	//fmt.Printf("row : %v\n", *user)
}

func set(db *sql.DB) {

}

func delete(db *sql.DB) {

}

func transaction() {

}

func beginTransaction() {

}
