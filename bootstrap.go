package main

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/bootstrap/",bootstrapHandler)
}

func bootstrapHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path;
	fmt.Fprintf(w,"%s",path)
}
