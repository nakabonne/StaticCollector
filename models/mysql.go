package models

import (
	"database/sql"
	"log"
)

var mysqlDB *sql.DB

func openMysql() error {
	var err error
	mysqlDB, err = sql.Open("mysql", "root:Tsuba2896@/web_crawler")
	return err
}
func closeMysql() {
	mysqlDB.Close()
	log.Println("MySQLの接続がCloseしました。")
}
