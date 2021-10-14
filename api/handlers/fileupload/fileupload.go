package handlers

import (
	"fmt"
	"log"
	"net/http"
	"tis-gf-api/api/fileupload"
)

type FileUploader struct {
	l *log.Logger
}

func NewFileUploader(l *log.Logger) *FileUploader {
	return &FileUploader{l}
}

func (acc *FileUploader) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	log.Fatalf("header: %v, body: %v", r.Header, r.Body)
	fmt.Printf("header: %v, body: %v", r.Header, r.Body)
	err = fileupload.Upload(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
