package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	http.HandleFunc("/files/",fileHandler)
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path,"/")
	possibleExtensions := strings.SplitAfter(r.URL.Path, ".")
	extension := possibleExtensions[len(possibleExtensions) - 1]
	switch(extension) {
	case "css":
		w.Header().Set("Content-Type", "text/css")
	case "js":
		w.Header().Set("Content-Type", "text/javascript")
	case "png":
		w.Header().Set("Content-Type", "image/png")
	}
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(404)
		return
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	fileSize := fileInfo.Size()
	fileData := make([]byte, fileSize, fileSize)
	numBytesRead, err := file.Read(fileData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	if int64(numBytesRead) != fileSize {
		log.Printf("Expected %v bytes, read %v bytes\n", fileSize, numBytesRead)
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w,"%s", string(fileData))
}
