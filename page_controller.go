package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"webCrawler/models"
)

type searchPages struct {
	StaticFiles []*models.StaticFile
	Pages       []*models.Page
	Keywords    []*models.Keyword
	PageID      int
	KeywordID   int
}

func pageSearch(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/search.tmpl"))
	if err := temp.Execute(w, &searchPages{
		Pages:    models.AllPages(mysqlDB),
		Keywords: models.AllKeywords(mysqlDB),
	}); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageCompetitorIndex(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	pageID, _ := strconv.Atoi(strings.Join(r.Form["page_id"], ""))
	keywordID, _ := strconv.Atoi(strings.Join(r.Form["keyword_id"], ""))
	fmt.Println(pageID, keywordID)

	staticFiles := models.FindStaticFilesByPageWord(pageID, keywordID, mongoDB)
	/*days := make([]string, len(staticFiles))
	for _, v := range staticFiles {
		days = append(days, v.TargetDay.Format("2006/01/02 Mon"))
	}*/

	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/search.tmpl"))
	if err := temp.Execute(w, &searchPages{
		StaticFiles: staticFiles,
		Pages:       models.AllPages(mysqlDB),
		Keywords:    models.AllKeywords(mysqlDB),
		PageID:      pageID,
		KeywordID:   keywordID,
	}); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageComparison(w http.ResponseWriter, r *http.Request) {
	// リクエストをパース
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	for i, v := range r.Form["day_0"] {
		fmt.Println(i)
		fmt.Println(v)
	}

	pageID, _ := strconv.Atoi(strings.Join(r.Form["page_id"], ""))
	keywordID, _ := strconv.Atoi(strings.Join(r.Form["keyword_id"], ""))
	layout := "2006-01-02 15:04:05 -0700 MST"
	targetDay0, _ := time.Parse(layout, r.Form["days[]"][0])
	targetDay1, _ := time.Parse(layout, r.Form["days[]"][1])

	//hei := strings.Join(r.Form["page_id"], "")

	// TODO 仮のstaticFiles
	staticFiles := make([]models.StaticFile, 0)
	//staticFiles = append(staticFiles, *models.FindStaticFilesByPageWord(8, 1, mongoDB)[0], *models.FindStaticFilesByPageWord(8, 1, mongoDB)[0])
	staticFiles = append(staticFiles, *models.FindStaticFilesByPageWordTargetday(pageID, keywordID, targetDay0, mongoDB))
	staticFiles = append(staticFiles, *models.FindStaticFilesByPageWordTargetday(pageID, keywordID, targetDay1, mongoDB))
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/comparison.tmpl"))
	if err := temp.Execute(w, staticFiles); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
