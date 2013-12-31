package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func init() {
	http.HandleFunc("/bootstrap/",bootstrapHandler)
}

func bootstrapHandler(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path,"/")
	//possibleExtensions := strings.SplitAfter(r.URL.Path, ".")
	//extension := possibleExtensions[len(possibleExtensions) - 1]
	//switch(extension) {
	//case "css":
	//	
	//case "js":
	//	
	//}
	bsFile, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(404)
		return
	}
	bsFileInfo, err := bsFile.Stat()
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	bsFileSize := bsFileInfo.Size()
	bsFileData := make([]byte, bsFileSize, bsFileSize)
	numBytesRead, err := bsFile.Read(bsFileData)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(500)
		return
	}
	if int64(numBytesRead) != bsFileSize {
		log.Printf("Expected %v bytes, read %v bytes\n", bsFileSize, numBytesRead)
		w.WriteHeader(500)
		return
	}
	bsMIMEType := http.DetectContentType(bsFileData)
	w.Header.Set("Content-Type", bsMIMEType)
	fmt.Fprintf(w,"%s", string(bsFileData))
}