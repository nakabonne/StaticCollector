package models

import (
	"log"
	"strconv"
)

// キーワードを管理
type Keyword struct {
	ID   int
	Word string
}

// Insert インサートする
func (k *Keyword) Insert() {
	query := "INSERT INTO keywords (id, word) values(?, ?)"
	if _, err := mysqlDB.Exec(query, k.ID, k.Word); err != nil {
		log.Fatal("インサートエラー：", err)
	}
}

func AllKeywords() []*Keyword {
	rows, err := mysqlDB.Query("SELECT * FROM `keywords`")
	if err != nil {
		log.Fatal("クエリーエラー：", err)
		// なんか返す
	}

	keywords := []*Keyword{}

	for rows.Next() {
		var (
			id   int
			word string
		)
		if err := rows.Scan(&id, &word); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		keywords = append(keywords, &Keyword{ID: id, Word: word})
	}
	rows.Close()
	return keywords
}

func FindKeyword(id int) *Keyword {

	query := "SELECT * FROM `keywords` WHERE `id` = " + strconv.Itoa(id)
	log.Println(query)
	rows, err := mysqlDB.Query(query)
	if err != nil {
		log.Fatal("クエリーエラー：", err)
	}

	var keyword *Keyword

	for rows.Next() {
		var (
			id   int
			word string
		)
		if err := rows.Scan(&id, &word); err != nil {
			log.Fatal("スキャンエラー: ", err)
		}
		keyword = &Keyword{ID: id, Word: word}
	}
	rows.Close()
	return keyword

}
