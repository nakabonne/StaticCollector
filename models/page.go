package models

import (
	"log"
	"strings"
)

type Page struct {
	ID  int
	URL string
}

// Insert インサートする
func (p *Page) Insert() (err error) {
	query := "INSERT INTO pages (id, url) values(?, ?)"
	if _, err = mysqlDB.Exec(query, p.ID, p.URL); err != nil {
		return
	}
	return
}

// AllPages Pagesテーブルから全件取得
func AllPages() (pages []*Page, err error) {
	rows, err := mysqlDB.Query("SELECT * FROM `pages`")
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id  int
			url string
		)
		if err = rows.Scan(&id, &url); err != nil {
			return
		}
		pages = append(pages, &Page{ID: id, URL: url})
	}
	rows.Close()
	return
}

func FormatURL(u string) string {
	str := strings.Replace(u, "https://www.google.co.jp/url?q=", "", 1)
	URLarray := strings.Split(str, "&")
	return URLarray[0]
}

func FindPageByURL(u string) (page *Page, err error) {
	query := "SELECT * FROM `pages` WHERE `url` = '" + u + "'"
	log.Println(query)
	rows, err := mysqlDB.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id  int
			url string
		)
		if err = rows.Scan(&id, &url); err != nil {
			return
		}
		page = &Page{ID: id, URL: url}
	}
	rows.Close()
	return
}
