package models

import (
	"database/sql"
	"log"
)

// OpenMysql Mysqlのコネクションを開く
func OpenMysql() *sql.DB {
	db, err := sql.Open("mysql", "root:Tsuba2896@/web_crawler")
	if err != nil {
		log.Fatal("エラー：", err)
	}
	return db
}
