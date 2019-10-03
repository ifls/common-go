package orm

import (
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
	"gocore/util"
)

const (
	DBTYPE_USER   = "user"
	DBTYPE_CONFIG = "config"
	DBTYPE_FILE   = "file"
	DBTYPE_HOT    = "hot"
	DBTYPE_LOG    = "log"
	DBTYPE_STAT   = "stat"
)

var dbs map[string]*xorm.Engine

var currentDb *xorm.Engine

func init() {
	dbs = make(map[string]*xorm.Engine)

	db, err := connectDb("")
	//db, err := connectDb(db.MysqlUrl("root", "02#F20ebac", "tcp", host, 3306, "test_db"))
	util.LogErrAndExit(err)
	dbs["default"] = db
	currentDb = db
}

func connectDb(url string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine("mysql", url)
	if err != nil {
		util.LogErr(err, zap.String("reason", "new engine connect"))
		return nil, err
	}
	return db, err
}

//根据表结构, 创建对应的表
func CreateTable(tableStruct interface{}) error {
	return createTable(currentDb, tableStruct)
}

//根据表结构, 创建对应的表
func createTable(db *xorm.Engine, tableStruct interface{}) error {
	if err := db.CreateTables(tableStruct); err != nil {
		util.LogErr(err, zap.String("reason", "create table"))
		return err
	}
	return nil
}

func AlterTable(tablename interface{}, tableStruct interface{}) error {
	return nil
}

func alterTable(db interface{}, tablename interface{}, tableStruct interface{}) error {
	return nil
}

func DeleteTable(tableStruct interface{}) error {
	return deleteTable(currentDb, tableStruct)
}

func deleteTable(db *xorm.Engine, tableStruct interface{}) error {
	err := db.DropTables(tableStruct)
	if err != nil {
		util.LogErr(err, zap.String(util.LOGTAG_REASON, "delete table error"))
		return err
	}
	return nil
}

func GetTableInfo(db interface{}, table interface{}) (interface{}, error) {
	return nil, nil
}

func Find(beanSlice interface{}, query interface{}, args ...interface{}) error {
	return find(currentDb, beanSlice, query, args...)
}

func find(db *xorm.Engine, beanSlice interface{}, query interface{}, args ...interface{}) error {
	err := db.Where(query, args...).Find(beanSlice)
	if err != nil {
		util.LogErr(err)
		return err
	}

	return nil
}

func Get(bean interface{}, query interface{}, args ...interface{}) (bool, error) {
	return get(currentDb, bean, query, args...)
}

func get(db *xorm.Engine, bean interface{}, query interface{}, args ...interface{}) (bool, error) {
	has, err := db.Where(query, args...).Get(bean)
	if err != nil {
		util.LogErr(err)
		return false, err
	}

	if has {
		return true, nil
	}

	return false, nil
}

func UpdateRecord(db interface{}, table interface{}, where interface{}) error {
	return nil
}

func InsertRecord(beans ...interface{}) error {
	return insertRecord(currentDb, beans...)
}

func insertRecord(db *xorm.Engine, beans ...interface{}) error {
	affected, err := db.Insert(beans...)
	if err != nil {
		util.LogErr(err)
		return err
	}

	if int(affected) != len(beans) {
		errStr := fmt.Sprintf("insert %d row, but %d row success", len(beans), affected)
		util.LogError(errStr)
		return errors.New(errStr)
	}
	return nil
}

func Exec(sql string) ([]map[string][]byte, error) {
	return exec(currentDb, sql)
}

func exec(db *xorm.Engine, sql string) ([]map[string][]byte, error) {
	results, err := db.Query(sql)
	if err != nil {
		util.LogErr(err)
		return nil, err
	}
	return results, nil
}

func DeleteRecord(db interface{}, table interface{}, where interface{}) error {
	return nil
}

func CreateDB(dbname string) {

}

func DeleteDB(db interface{}) {

}

func getInfo() {

}

func UseDB(dbType string) {

}
