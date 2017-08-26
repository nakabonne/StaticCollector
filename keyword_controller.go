package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"webCrawler/models"
)

func keywordInsert(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/keyword/insert.tmpl"))
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
	keyword := models.Keywords{
		ID:   4,
		Word: word,
	}
	keyword.Insert()
	http.Redirect(w, r, "/keyword/insert", 301)
}
