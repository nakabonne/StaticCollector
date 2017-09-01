package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"webCrawler/models"
)

func keywordInsert(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/keyword/insert.tmpl"))
	if err := temp.Execute(w, nil); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func keywordCreate(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	word := strings.Join(r.Form["word"], "")
	keyword := &models.Keywords{Word: word}
	keyword.Insert(mysqlDB)
	http.Redirect(w, r, "/keyword/insert", 301)
}

func keywordCrawl(w http.ResponseWriter, r *http.Request) {
	keywords := models.AllKeywords(mysqlDB)
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/keyword/crawl.tmpl"))
	if err := temp.Execute(w, keywords); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func crawl(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	word := strings.Join(r.Form["keyword"], "")

	log.Println("検索ワード：", word)
	word = strings.Replace(word, " ", "+", -1)
	firstURL := "https://www.google.co.jp/search?rlz=1C5CHFA_enJP693JP693&q=" + string(word)
	log.Println("検索URL：", firstURL)

	c := newCrawler()
	go c.collectHTML()
	wordID := 1 // SQLから取得する
	c.req <- &request{
		url:    firstURL,
		wordID: wordID,
		depth:  2,
	}

	http.Redirect(w, r, "/keyword/crawl", 301)
}
