package models

import (
	"database/sql"
	"log"
)

type Pages struct {
	ID    int
	URL   string
	Title string
}

// Insert インサートする
func (p *Pages) Insert(db *sql.DB) {
	query := "INSERT INTO pages (id, url, title) values(?, ?, ?)"
	if _, err := db.Exec(query, p.ID, p.URL, p.Title); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}

func FindPageByURL(db *sql.DB, u string) *Pages {
	query := "SELECT * FROM `pages` WHERE `url` = " + u
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("クエリーエラー：", err)
	}

	var page *Pages

	for rows.Next() {
		var (
			id  int
			url string
		)
		if err := rows.Scan(&id, &url); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		page = &Pages{ID: id, URL: url}
	}
	rows.Close()
	return page
}
