package main

import (
	"html/template"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", homePageHandler)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
}