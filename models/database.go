package models

import (
	"database/sql"
	"log"

	mgo "gopkg.in/mgo.v2"
)

// OpenMysql Mysqlのコネクションを開く
func OpenMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:Tsuba2896@/web_crawler")
	if err != nil {
		log.Fatal("Mysqlオープン時にエラー：", err)
	}
	return db
}

func GetSettionMongo() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost/")
	if err != nil {
		log.Fatal("mongoDBオープン時にエラー", err)
	}
	return session
}
