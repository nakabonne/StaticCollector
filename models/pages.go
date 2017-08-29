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
