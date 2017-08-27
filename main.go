package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	"webCrawler/models"

	_ "github.com/go-sql-driver/mysql"
)

var mysqlDB = models.OpenMysql()
var mongoDB = models.GetSettionMongo()

func main() {
	// いずれ消す---------------------------------------------------------------------
	var word = flag.String("w", " ", "検索ワードを入力して下さい")
	flag.Parse()
	log.Println("検索ワード：", *word)
	*word = strings.Replace(*word, " ", "+", -1)
	firstURL := "https://www.google.co.jp/search?rlz=1C5CHFA_enJP693JP693&q=" + string(*word)
	log.Println("検索URL：", firstURL)
	// -------------------------------------------------------------------------------

	defer mysqlDB.Close()
	defer mongoDB.Clone()
	// クローリング開始-------------
	c := newCrawler()
	go c.collectHTML()
	c.req <- &request{
		url:   firstURL,
		depth: 2,
	}
	// -----------------------------

	http.HandleFunc("/keyword/insert", keywordInsert)
	http.HandleFunc("/keyword/create", keywordCreate)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
