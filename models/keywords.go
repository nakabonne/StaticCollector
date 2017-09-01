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

func AllKeywords(db *sql.DB) []*Keywords {
	rows, err := db.Query("SELECT * FROM `keywords`")
	if err != nil {
		log.Fatal("クエリーエラー：", err)
		// なんか返す
	}

	keywords := []*Keywords{}

	for rows.Next() {
		var (
			id   int
			word string
		)
		if err := rows.Scan(&id, &word); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		keywords = append(keywords, &Keywords{ID: id, Word: word})
	}
	rows.Close()
	return keywords
}
