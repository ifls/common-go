package mysql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ifls/gocore/util"
	"time"
)

const (
	TABLE_USER = "create table proto(uid int unique key auto_increment, password varchar(20), ID varchar(18));"
)

func MysqlUrl(user, password, host string, port uint16, database string) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, "tcp", host, port, database)
}

var testClient *sql.DB
var defaultClient *sql.DB

func init() {
	host := ""
	port := 99999
	url := MysqlUrl("root", "02#F20ebac", host, uint16(port), "test")
	testClient = open(url)
	defaultClient = testClient
	Ping()
}

type MysqlUser struct {
	uid      int64  `db:"uid"`
	ID       string `db:"ID"` //由于在mysql的users表中name没有设置为NOT NULL,所以name可能为null,在查询过程中会返回nil，如果是string类型则无法接收nil,但sql.NullString则可以接收nil值
	password string `db:"password"`
}

func defaultSetting(mysql *sql.DB) {
	mysql.SetConnMaxLifetime(10 * time.Second) //最大连接周期，超过时间的连接就close, 过期的连接在下次使用时才会被lazyclose,
	mysql.SetMaxOpenConns(100)                 //设置最大连接数
	mysql.SetMaxIdleConns(16)                  //设置闲置连接赤的最大值
}

//测试数据库连通性
func Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	err := defaultClient.PingContext(ctx)
	if err != nil {
		util.LogErr(err)
		return err
	}
	return nil
}

func open(url string) (db *sql.DB) {
	db, err := sql.Open("mysql", url)
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
		util.LogErr(err)
		return err
	}

	//获取插入数据的主键id,判断是否插入成功
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		util.LogErr(err)
		return err
	}

	if lastInsertID < 1 {
		err := fmt.Errorf("%s", "insert id < 1")
		util.LogErr(err)
		return err
	}

	util.DevInfo("LastInsertID:", lastInsertID)
	rowsAffected, err := result.RowsAffected() //影响行数
	if err != nil {
		util.LogErr(err)
		return err
	}

	if rowsAffected < 1 {
		err := errors.New("insert affect zero row")
		util.LogErr(err)
		return err
	}

	util.DevInfo("RowsAffected:", rowsAffected)
	return nil
}

//返回一个sql语句,以后执行
func _prepare(db *sql.DB) {

}

//queryRow执行一个sql, 返回至多一行
func query(db *sql.DB, sql string) error {
	user := new(MysqlUser)
	row := db.QueryRow("select * from proto where uid = 3")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.uid, &user.ID, &user.password); err != nil {
		util.LogErr(err)
		return err
	}
	util.DevInfo("row : %v\n", *user)
	return nil
}

func queryWithTimeOut(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := new(MysqlUser)
	row := db.QueryRowContext(ctx, "select * from proto where uid = 3")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	if err := row.Scan(&user.uid, &user.ID, &user.password); err != nil {
		fmt.Printf("scan failed, err:%v", err)
		return
	}
	fmt.Printf("row : %v\n", *user)
}

func set(db *sql.DB) {

}

func delete(db *sql.DB) {

}

/// 事务 ///
func transaction() {

}

func beginTransaction() {

}
