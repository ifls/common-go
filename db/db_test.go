package db

import "testing"

func TestDb(t *testing.T) {
	//TestRedis()
	//TestMongo()
	//TestMysql()
	InitDB("")
	GetDB("", "")
}
