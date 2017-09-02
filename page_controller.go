package main

import (
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
	keywords, err := models.AllKeywords()
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/keyword/insert", 301)
	}
	pages, err := models.AllPages()
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/keyword/insert", 301)
	}
	if err := temp.Execute(w, &searchPages{
		Pages:    pages,
		Keywords: keywords,
	}); err != nil {
		log.Fatal("template error!!!!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageCompetitorIndex(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/page/serach", 301)
	}

	pageID, _ := strconv.Atoi(strings.Join(r.Form["page_id"], ""))
	keywordID, _ := strconv.Atoi(strings.Join(r.Form["keyword_id"], ""))

	staticFiles, err := models.FindStaticFilesByPageWord(pageID, keywordID)
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/page/serach", 301)
	}
	keywords, err := models.AllKeywords()
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/page/serach", 301)
	}
	pages, err := models.AllPages()
	if err != nil {
		log.Fatal(err)
		http.Redirect(w, r, "/page/serach", 301)
	}
	searchPages := &searchPages{
		StaticFiles: staticFiles,
		Pages:       pages,
		Keywords:    keywords,
		PageID:      pageID,
		KeywordID:   keywordID,
	}
	sort.Sort(searchPages.StaticFiles)
	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/search.tmpl"))
	if err := temp.Execute(w, searchPages); err != nil {
		log.Fatal("template error", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageComparison(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal("parse error!!!!", err)
		http.Redirect(w, r, "/page/serach", 301)
	}

	pageID, _ := strconv.Atoi(strings.Join(r.Form["page_id"], ""))
	keywordID, _ := strconv.Atoi(strings.Join(r.Form["keyword_id"], ""))
	days := r.Form["days[]"]

	//layout := "2006-01-02 15:04:05 -0700 MST"
	layout := "2006-01-02T15:04:05-07:00"
	var staticFiles []models.StaticFile
	if len(days) < 1 {
		http.Redirect(w, r, "/page/search", 301)
		return
	}
	if len(days) >= 1 {
		beforeDay, _ := time.Parse(layout, days[0])
		beforeStaticFile, err := models.FindStaticFilesByPageWordTargetday(pageID, keywordID, beforeDay)
		if err != nil {
			log.Fatal(err)
			http.Redirect(w, r, "/page/search", 301)
			return
		}
		staticFiles = append(staticFiles, *beforeStaticFile)
	}
	if len(days) >= 2 {
		afterDay, _ := time.Parse(layout, days[1])
		afterbeforeStaticFile, err := models.FindStaticFilesByPageWordTargetday(pageID, keywordID, afterDay)
		if err != nil {
			log.Fatal(err)
			http.Redirect(w, r, "/page/search", 301)
			return
		}
		staticFiles = append(staticFiles, *afterbeforeStaticFile)
	}

	temp := template.Must(template.ParseFiles("views/layout.tmpl", "views/page/comparison.tmpl"))
	if err := temp.Execute(w, staticFiles); err != nil {
		log.Fatal("template error!!!!", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
