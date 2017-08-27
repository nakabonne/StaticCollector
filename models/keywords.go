package models

import (
	"database/sql"
	"log"
)

// キーワードを管理
type Keywords struct {
	ID   int
	Word string
}

// Insert インサートする
func (k *Keywords) Insert(db *sql.DB) {
	query := "INSERT INTO keywords (id, word) values(?, ?)"
	if _, err := db.Exec(query, k.ID, k.Word); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}
