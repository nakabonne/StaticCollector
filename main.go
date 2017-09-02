package main

import (
	"log"
	"net/http"
	"webCrawler/models"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	models.OpenDB()
	defer models.CloseDB()

	//http.Handle("/lib/assets/", http.StripPrefix("/lib/assets/", http.FileServer(http.Dir("lib/assets/"))))
	http.HandleFunc("/keyword/insert", keywordInsert)
	http.HandleFunc("/keyword/create", keywordCreate)
	http.HandleFunc("/keyword/crawl", keywordCrawl)
	http.HandleFunc("/crawl", crawl)
	http.HandleFunc("/page/search", pageSearch)
	http.HandleFunc("/page/competitor", pageCompetitorIndex)
	http.HandleFunc("/page/comparison", pageComparison)

	log.Println("準備OK")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndSearver:", err)
	}
}
