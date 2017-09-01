package models

import (
	"database/sql"
	"log"
	"strings"
)

type Pages struct {
	ID  int
	URL string
}

// Insert インサートする
func (p *Pages) Insert(db *sql.DB) {
	query := "INSERT INTO pages (id, url) values(?, ?)"
	if _, err := db.Exec(query, p.ID, p.URL); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}

// AllPages Pagesテーブルから全件取得
func AllPages(db *sql.DB) []*Pages {
	rows, err := db.Query("SELECT * FROM `pages`")
	if err != nil {
		log.Fatal("クエリーエラー：", err)
		// なんか返す
	}

	pages := []*Pages{}

	for rows.Next() {
		var (
			id  int
			url string
		)
		if err := rows.Scan(&id, &url); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		pages = append(pages, &Pages{ID: id, URL: url})
	}
	rows.Close()
	return pages
}

func FormatURL(u string) string {
	str := strings.Replace(u, "https://www.google.co.jp/url?q=", "", 1)
	URLarray := strings.Split(str, "&")
	return URLarray[0]
}

func FindPageByURL(db *sql.DB, u string) *Pages {
	query := "SELECT * FROM `pages` WHERE `url` = '" + u + "'"
	log.Println(query)
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
