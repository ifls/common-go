package db

import "log"

func init() {

}

func InitDB(conf string) {
	log.Println(conf)
}

func GetDB(dbType string, database string) {
	log.Println(dbType, database)
}
