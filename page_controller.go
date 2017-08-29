package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func pageSearch(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/page/search.tmpl"))
	if err := temp.Execute(w, nil); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageCompetitorIndex(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	page_url := strings.Join(r.Form["page_url"], "")
	fmt.Print(page_url)

}
