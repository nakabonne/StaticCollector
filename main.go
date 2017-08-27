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
	defer mysqlDB.Close()
	defer mongoDB.Clone()
	var word = flag.String("w", " ", "検索ワードを入力して下さい")
	flag.Parse()
	log.Println("検索ワード：", *word)
	*word = strings.Replace(*word, " ", "+", -1)
	firstURL := "https://www.google.co.jp/search?rlz=1C5CHFA_enJP693JP693&q=" + string(*word)
	log.Println("検索URL：", firstURL)
	m := newCrawler()
	go m.execute()
	m.req <- &request{
		url:   firstURL,
		depth: 2,
	}

	/* mongoインサート方法--------------------------------
		page := &models.Pages{
			ID:        bson.NewObjectId(),
			Title:     "行くぜ",
			URL:       "iku.com",
			HTML:      "<html></html>",
			Rank:      1,
			TargetDay: time.Now(),
		}
		page.Insert(mongoDB)
	---------------------------------------------------*/
	http.HandleFunc("/keyword/insert", keywordInsert)
	http.HandleFunc("/keyword/create", keywordCreate)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
