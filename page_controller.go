package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"webCrawler/models"
)

type searchPages struct {
	StaticFiles models.StaticFiles
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
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	pageID, _ := strconv.Atoi(strings.Join(r.Form["page_id"], ""))
	keywordID, _ := strconv.Atoi(strings.Join(r.Form["keyword_id"], ""))

	staticFiles := models.FindStaticFilesByPageWord(pageID, keywordID, mongoDB)
	searchPages := &searchPages{
		StaticFiles: staticFiles,
		Pages:       models.AllPages(mysqlDB),
		Keywords:    models.AllKeywords(mysqlDB),
		PageID:      pageID,
		KeywordID:   keywordID,
	}
	sort.Sort(searchPages.StaticFiles)
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/search.tmpl"))
	if err := temp.Execute(w, searchPages); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageComparison(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("エラー：", err)
	}

	pageID, _ := strconv.Atoi(strings.Join(r.Form["page_id"], ""))
	keywordID, _ := strconv.Atoi(strings.Join(r.Form["keyword_id"], ""))
	days := r.Form["days[]"]
	fmt.Println("daysは", days)

	layout := "2006-01-02 15:04:05 -0700 MST"
	var staticFiles []models.StaticFile
	if len(days) < 1 {
		fmt.Println("0つ")
		http.Redirect(w, r, "/page/competitor", 301)
		return
	}
	if len(days) >= 1 {
		fmt.Println("1つ")
		beforeDay, _ := time.Parse(layout, days[0])
		fmt.Println("いえ")
		beforeStaticFile, err := models.FindStaticFilesByPageWordTargetday(pageID, keywordID, beforeDay, mongoDB)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("いお")
		staticFiles = append(staticFiles, *beforeStaticFile)
		fmt.Println("1つ完了")
	}
	if len(days) >= 2 {
		fmt.Println("2つ")
		afterDay, _ := time.Parse(layout, days[1])
		afterbeforeStaticFile, err := models.FindStaticFilesByPageWordTargetday(pageID, keywordID, afterDay, mongoDB)
		if err != nil {
			log.Fatal(err)
		}
		staticFiles = append(staticFiles, *afterbeforeStaticFile)
	}

	fmt.Println("スタティックは", staticFiles)

	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/comparison.tmpl"))
	if err := temp.Execute(w, staticFiles); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
