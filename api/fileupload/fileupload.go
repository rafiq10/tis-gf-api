package fileupload

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Origin", "*")
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers:", "*")
	w.Header().Set("Content-Type", "multipart/form-data, application/x-www-form-urlencoded, text/json")

	fmt.Println("File Upload init ...")
	log.Fatal(r.Header, r.Body)
	r.ParseMultipartForm(1 << 2)

	f, h, err := r.FormFile("fileName")

	if err != nil {
		fmt.Fprintf(w, "error retrieving file\n %v", h.Filename)
		fmt.Printf(err.Error())
		return err
	}

	defer f.Close()
	fmt.Printf("Uploaded file: %+v\n", h.Filename)
	fmt.Printf("File size: %+v\n", h.Size)
	fmt.Printf("Mime header: %+v\n", h.Header)

	tmpF, err := ioutil.TempFile("/home/bilrafal/Documents", "tmp-")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer tmpF.Close()

	fBytes, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	tmpF.Write(fBytes)
	fmt.Fprintf(w, "Successfully uploaded file:\n %v", h.Filename)
	return nil
}
