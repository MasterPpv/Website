package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("Server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}