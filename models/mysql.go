package models

import (
	"database/sql"
	"log"
)

var mysqlDB *sql.DB

func openMysql() (err error) {
	// sqlUser,sqlPass,sqlNameは別ファイルにて管理(gitでは管理外)
	mysqlDB, err = sql.Open("mysql", sqlUser+":"+sqlPass+"@/"+sqlName)
	return
}
func closeMysql() {
	mysqlDB.Close()
	log.Println("MySQLの接続がCloseしました。")
}
