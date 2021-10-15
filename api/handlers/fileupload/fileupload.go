package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type FileUploader struct {
	l *log.Logger
}

func NewFileUploader(l *log.Logger) *FileUploader {
	return &FileUploader{l}
}

func (fu *FileUploader) FileUpload(w http.ResponseWriter, r *http.Request) {

	r.ParseMultipartForm(1 << 2)

	f, h, err := r.FormFile("fileName")

	if err != nil {
		http.Error(w, "error retrieving file: '"+h.Filename+"'\n "+err.Error(), http.StatusInternalServerError)
	}

	defer f.Close()
	fmt.Printf("Uploaded file: %+v\n", h.Filename)
	fmt.Printf("File size: %+v\n", h.Size)
	fmt.Printf("Mime header: %+v\n", h.Header)

	tmpF, err := ioutil.TempFile("/home/bilrafal/Documents", "tmp-")

	if err != nil {
		http.Error(w, "Error creating tmp file.\n "+err.Error(), http.StatusInternalServerError)
	}

	defer tmpF.Close()

	fBytes, err := ioutil.ReadAll(f)

	if err != nil {
		http.Error(w, "Error reading file.\n "+err.Error(), http.StatusInternalServerError)
	}

	tmpF.Write(fBytes)
	fmt.Fprintf(w, "Successfully uploaded file:\n %v", h.Filename)
}
