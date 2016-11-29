package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"io"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./gtpls/handleUpload.gtpl")
		t.Execute(w, getToken())

	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadFile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
	//	filename := handler.Filename
		filename := getFileName(handler.Filename)
		fmt.Printf(filename)
		f, err := os.OpenFile("./uploaded/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

	}
}
