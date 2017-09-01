package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
	"webCrawler/models"
)

type searchPages struct {
	StaticFiles []*models.StaticFiles
	Days        []string
	Pages       []*models.Pages
	Keywords    []*models.Keywords
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
	days := make([]string, len(staticFiles))
	for _, v := range staticFiles {
		days = append(days, v.TargetDay.Format("2006/01/02 Mon"))
	}

	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/search.tmpl"))
	if err := temp.Execute(w, &searchPages{
		StaticFiles: staticFiles,
		Days:        days,
		Pages:       models.AllPages(mysqlDB),
		Keywords:    models.AllKeywords(mysqlDB),
	}); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageComparison(w http.ResponseWriter, r *http.Request) {
	// TODO 仮のstaticFiles
	staticFiles := make([]models.StaticFiles, 2)
	staticFiles = append(staticFiles, *models.FindStaticFilesByPageWord(24, 1, mongoDB)[0], *models.FindStaticFilesByPageWord(24, 1, mongoDB)[0])
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/comparison.tmpl"))
	if err := temp.Execute(w, staticFiles); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
