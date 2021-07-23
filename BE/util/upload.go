package util

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := handler.Filename
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	f, err := os.OpenFile("./public/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, file)
	json.NewEncoder(w).Encode(map[string]string{
		"url": "http://localhost:8000/public/" + fileName,
	})
}
