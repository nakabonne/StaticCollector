package models

import (
	"database/sql"
	"log"

	mgo "gopkg.in/mgo.v2"
)

var mysqlDB *sql.DB
var mongoSession *mgo.Session
var mongoDB *mgo.Database

// OpenMysql Mysqlのコネクションを開く
func OpenMysql() error {
	var err error
	mysqlDB, err = sql.Open("mysql", "root:Tsuba2896@/web_crawler")
	return err
}
func CloseMysql() {
	mysqlDB.Close()
	log.Println("MySQLの接続がCloseしました。")
}

func DialMongo() error {
	var err error
	mongoSession, err = mgo.Dial("mongodb://localhost/")
	return err
}
func SetMongoDB() {
	mongoDB = mongoSession.DB("web_crawler")
}
func CloseMongo() {
	mongoSession.Clone()
	log.Println("mongoDB接続がCloseしました。")
}

func getCollection(name string) *mgo.Collection {
	return mongoDB.C(name)
}
