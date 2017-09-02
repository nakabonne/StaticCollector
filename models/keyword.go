package models

import (
	"database/sql"
	"log"
	"strconv"
)

// キーワードを管理
type Keyword struct {
	ID   int
	Word string
}

// Insert インサートする
func (k *Keyword) Insert() (err error) {
	query := "INSERT INTO keywords (id, word) values(?, ?)"
	if _, err = mysqlDB.Exec(query, k.ID, k.Word); err != nil {
		return
	}
	return
}

func AllKeywords() (keywords []*Keyword, err error) {
	var rows *sql.Rows
	query := "SELECT * FROM `keywords`"
	rows, err = mysqlDB.Query(query)
	log.Println(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id   int
			word string
		)
		if err = rows.Scan(&id, &word); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		keywords = append(keywords, &Keyword{ID: id, Word: word})
	}
	rows.Close()
	return
}

func FindKeyword(id int) (keyword *Keyword, err error) {
	query := "SELECT * FROM `keywords` WHERE `id` = " + strconv.Itoa(id)
	var rows *sql.Rows
	rows, err = mysqlDB.Query(query)
	log.Println(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var (
			id   int
			word string
		)
		if err = rows.Scan(&id, &word); err != nil {
			return
		}
		keyword = &Keyword{ID: id, Word: word}
	}
	rows.Close()
	return
}
