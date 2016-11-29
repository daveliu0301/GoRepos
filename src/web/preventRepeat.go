package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func preventRepeat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./gtpls/preventRepeat.gtpl")
		t.Execute(w, getToken()) // second arg?
	} else {
		r.ParseForm()
		fmt.Fprintln(w, r.Form["input"][0])
		fmt.Fprintln(w, r.Form.Get("token"))
		fmt.Println(r.Form["input"][0])
		fmt.Println(r.Form.Get("token")) // different and usage r.Form.Get() & r.Form[][]
	}

}
