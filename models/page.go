package models

import (
	"database/sql"
	"log"
	"strings"
)

type Page struct {
	ID  int
	URL string
}

// Insert インサートする
func (p *Page) Insert(db *sql.DB) {
	query := "INSERT INTO pages (id, url) values(?, ?)"
	if _, err := db.Exec(query, p.ID, p.URL); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}

// AllPages Pagesテーブルから全件取得
func AllPages(db *sql.DB) []*Page {
	rows, err := db.Query("SELECT * FROM `pages`")
	if err != nil {
		log.Fatal("クエリーエラー：", err)
		// なんか返す
	}

	pages := []*Page{}

	for rows.Next() {
		var (
			id  int
			url string
		)
		if err := rows.Scan(&id, &url); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		pages = append(pages, &Page{ID: id, URL: url})
	}
	rows.Close()
	return pages
}

func FormatURL(u string) string {
	str := strings.Replace(u, "https://www.google.co.jp/url?q=", "", 1)
	URLarray := strings.Split(str, "&")
	return URLarray[0]
}

func FindPageByURL(db *sql.DB, u string) *Page {
	query := "SELECT * FROM `pages` WHERE `url` = '" + u + "'"
	log.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("クエリーエラー：", err)
	}

	var page *Page

	for rows.Next() {
		var (
			id  int
			url string
		)
		if err := rows.Scan(&id, &url); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		page = &Page{ID: id, URL: url}
	}
	rows.Close()
	return page
}
