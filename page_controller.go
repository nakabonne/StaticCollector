package main

import (
	"html/template"
	"log"
	"net/http"
)

func pageSearch(w http.ResponseWriter, r *http.Request) {
	temp := template.Must(template.ParseFiles("views/page/search.tmpl"))
	if err := temp.Execute(w, nil); err != nil {
		log.Fatal("テンプレートエラー", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func pageCompetitorIndex(w http.ResponseWriter, r *http.Request) {

}
