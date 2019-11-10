package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

type SqlCli struct {
	*xorm.Engine
	driver string
	url    string
}

func MakeMysqlUrl(user, password, host string, port uint16, database string) string {
	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8", user, password, "tcp", host, port, database)
}

//Open 打开数据库连接
// @driver string 数据库驱动类型
// @url string 类型对应的连接url
func Open(driver string, url string) (SqlCli, error) {
	db, err := xorm.NewEngine(driver, url)
	log.Println(url)
	if err != nil {
		log.Println(driver, url)
		return SqlCli{}, fmt.Errorf("open db fail, %s, %s %w", driver, url, err)
	}

	return SqlCli{
		Engine: db,
		driver: driver,
		url:    url,
	}, nil
}

// Close 关闭连接
func (cli SqlCli) Close() {

}

// 应该是运维的工作
func (cli SqlCli) CreateDb() {

}

// -------------------------common sql exec ---------------------------
//执行sql 查询语句
func (cli SqlCli) Query(sql string) ([]map[string]interface{}, error) {
	rows, err := cli.Engine.QueryInterface(sql)
	if err != nil {
		return nil, fmt.Errorf("query fail %s %w", sql, err)
	}

	return rows, nil
}

//执行sql update语句
func (cli SqlCli) Exec(sql string) error {
	result, err := cli.Engine.Exec(sql)
	if err != nil {
		return fmt.Errorf("exec fail %s %w", sql, err)
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("exec RowsAffected fail %w", err)
	}

	last, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("exec LastInsertId fail %w", err)
	}

	log.Printf("LastInsertId =%d, RowsAffected = %d\n", last, affected)
	return nil
}

//------------------------表内 api ---------------------

func (cli SqlCli) Insert(beans ...interface{}) error {
	affacted, err := cli.Engine.Insert(beans...)
	if err != nil {
		return fmt.Errorf("insert fail %w", err)
	}

	if affacted != int64(len(beans)) {
		return fmt.Errorf("insert want insert %d, but inserted %d", len(beans), affacted)
	}

	return nil
}

//获取一行数据， 放到bean里
func (cli SqlCli) Get(bean interface{}) error {
	//指定的值， 必须匹配
	has, err := cli.Engine.Get(bean)
	if err != nil {
		return fmt.Errorf("get err: %w", err)
	}

	log.Printf("has = %v\n", has)
	return nil
}

//只判断指定行存在，不返回数据
func (cli SqlCli) Exist(tableScheme interface{}) error {
	//指定的值， 必须匹配
	has, err := cli.Engine.Exist(tableScheme)
	if err != nil {
		return fmt.Errorf("exist err: %w", err)
	}

	log.Printf("has = %v\n", has)
	return nil
}

//Find beans 是作为判断类型 以及作为返回值的
func (cli SqlCli) Find(beans interface{}) error {
	//指定的值， 必须匹配
	err := cli.Engine.Find(beans)
	if err != nil {
		return fmt.Errorf("find err: %w", err)
	}

	log.Printf("beans = %#v\n", beans)
	return nil
}

//一行一行迭代，
func (cli SqlCli) Iterate(bean interface{}, callback func(idx int, bean interface{}) error) error {
	err := cli.Engine.Iterate(bean, callback)
	if err != nil {
		return fmt.Errorf("iterate err: %w", err)
	}

	return nil
}

func (cli SqlCli) Update(bean interface{}, id int64) error {
	affacted, err := cli.Engine.ID(id).Update(bean)
	if err != nil {
		return fmt.Errorf("update err: %w", err)
	}

	log.Printf("affacted = %#v\n", affacted)
	return nil
}

func (cli SqlCli) Delete(bean interface{}, id int64) error {
	affacted, err := cli.Engine.ID(id).Delete(bean)
	if err != nil {
		return fmt.Errorf("delete err: %w", err)
	}

	log.Printf("affacted = %#v\n", affacted)
	return nil
}

//统计行数
func (cli SqlCli) Count(bean interface{}) (int64, error) {
	return cli.Engine.Count(bean)
}

// 指定字段求和
func (cli SqlCli) Sum(bean interface{}, field string) (float64, error) {
	return cli.Engine.Sum(bean, field)
}

func (cli SqlCli) BeginTxn() *xorm.Session {
	return cli.Engine.NewSession()
}

func (cli SqlCli) EndTxn(session *xorm.Session) error {
	defer session.Close()
	return session.Commit()
}

func (cli SqlCli) Txn(actions func(session *xorm.Session) (interface{}, error)) error {
	result, err := cli.Engine.Transaction(actions)
	if err != nil {
		return fmt.Errorf("txn err: %w", err)
	}

	log.Printf("txn result = %#v\n", result)
	return nil
}

//----------------表级 api----------------------------
//根据表结构, 创建对应的表, 只能加字段，不能改表结构，不插入数据
func (cli SqlCli) CreateTable(tableScheme interface{}) error {
	return cli.Engine.Sync2(tableScheme)
}

//
//func AlterTable(tablename interface{}, tableStruct interface{}) error {
//	return nil
//}

func (cli SqlCli) DeleteTable(table interface{}) error {
	err := cli.Engine.DropTables(table)
	if err != nil {
		return fmt.Errorf("delete table :%w", err)
	}

	return nil
}

//读取表结构
//func (cli SqlCli) GetTableInfo(table interface{}) (interface{}, error) {
//
//}
