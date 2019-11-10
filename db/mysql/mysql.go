package mysql

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type MyMysqlCli struct {
	*sql.DB
	mysqlUrl string
}

func (cli MyMysqlCli) defaultSetting() {
	cli.SetConnMaxLifetime(10 * time.Second) //最大连接周期，超过时间的连接就close, 过期的连接在下次使用时才会被lazyclose,
	cli.SetMaxOpenConns(100)                 //设置最大连接数
	cli.SetMaxIdleConns(16)                  //设置闲置连接赤的最大值
}

func MakeMysqlUrl(user, password, host string, port uint16, database string) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s", user, password, "tcp", host, port, database)
}

/////////////// mysql public api //////////////////

//测试数据库连通性
func (cli MyMysqlCli) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	err := cli.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("PingContext %w", err)
	}

	return nil
}

func Open(url string) (cli MyMysqlCli, err error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return MyMysqlCli{}, fmt.Errorf("open %w", err)
	}

	return MyMysqlCli{
		DB:       db,
		mysqlUrl: url,
	}, nil
}

func (cli MyMysqlCli) CloseDb(db *sql.DB) {
	_ = cli.Close()
}

func (cli MyMysqlCli) CreateTable(db *sql.DB) {

}

func (cli MyMysqlCli) AlterTable(db *sql.DB) {

}

func (cli MyMysqlCli) GetTableScheme(db *sql.DB) {

}

func (cli MyMysqlCli) DeleteTable(db *sql.DB) {

}

func (cli MyMysqlCli) Insert(db *sql.DB) error {
	//result, err := db.Exec("insert into proto(password, id) values(?,?)", "password1", "430223199509050711")
	//if err != nil {
	//	util.LogErr(err)
	//	return err
	//}
	//
	////获取插入数据的主键id,判断是否插入成功
	//lastInsertID, err := result.LastInsertId()
	//if err != nil {
	//	util.LogErr(err)
	//	return err
	//}
	//
	//if lastInsertID < 1 {
	//	err := fmt.Errorf("%s", "insert id < 1")
	//	util.LogErr(err)
	//	return err
	//}
	//
	//util.DevInfo("LastInsertID: %v", lastInsertID)
	//rowsAffected, err := result.RowsAffected() //影响行数
	//if err != nil {
	//	util.LogErr(err)
	//	return err
	//}
	//
	//if rowsAffected < 1 {
	//	err := errors.New("insert affect zero row")
	//	util.LogErr(err)
	//	return err
	//}
	//
	//util.DevInfo("RowsAffected: %v", rowsAffected)
	return nil
}

//返回一个sql语句,以后执行
//func _prepare(db *sql.DB) {
//
//}

//queryRow执行一个sql, 返回至多一行
func (cli MyMysqlCli) Query(db *sql.DB, sql string) error {
	//user := new(MysqlUser)
	//_ = db.QueryRow("select * from proto where uid = 3")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	//if err := row.Scan(&user.uid, &user.ID, &user.password); err != nil {
	//	util.LogErr(err)
	//	return err
	//}
	//util.DevInfo("row : %v\n", *user)
	return nil
}

func (cli MyMysqlCli) QueryWithTimeOut(db *sql.DB) {
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//user := new(MysqlUser)
	//_ = db.QueryRowContext(ctx, "select * from proto where uid = 3")
	//row.scan中的字段必须是按照数据库存入字段的顺序，否则报错
	//if err := row.Scan(&user.uid, &user.ID, &user.password); err != nil {
	//	fmt.Printf("scan failed, err:%v", err)
	//	return
	//}
	//fmt.Printf("row : %v\n", *user)
}

//func Exec(db *sql.DB) {
//
//}
//
//func Update(db *sql.DB) {
//
//}
//
//func Delete(db *sql.DB) {
//
//}

/// 事务 ///
//func transaction() {
//
//}
//
//func beginTransaction() {
//
//}
