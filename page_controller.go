package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
	"webCrawler/models"
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

	pageID := strings.Join(r.Form["page_id"], "")
	keywordID := strings.Join(r.Form["keyword_id"], "")
	fmt.Println(pageID, keywordID)

	staticFiles := make([]*models.StaticFiles, 0)
	// TODO ①mongoからFindする
	// ②日付順にView表示
	// ③日付2つ選んで次男viewに渡す
	// AdminLTE導入
	// chart.jsでグラフ
	// HTML比較
	staticFiles = append(staticFiles, &models.StaticFiles{
		TargetDay: time.Now(),
		Rank:      1,
	})
	//a := time.Date(2001, 5, 31, 0, 0, 0, 0, time.Local)

	temp := template.Must(template.ParseFiles("views/page/search.tmpl"))
	if err := temp.Execute(w, staticFiles); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
