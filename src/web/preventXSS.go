package main

import (
	"fmt"
	"html/template"
	"net/http"
	textTemplate "text/template"
)

func preventXSS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./gtpls/preventXSS.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		input := r.Form["input"][0]
		//output to server log.
		fmt.Println("input: ", template.HTMLEscapeString(input))

		//output same js as input to client, must contain other chars before and behind js txst
		t, _ := textTemplate.New("foo").Parse(`{{define "T"}}text/template: {{.}}!{{end}}`) // what is it?
		t.ExecuteTemplate(w, "T", input)

		//another way for output same js as input to client
		t1, _ := template.New("foo").Parse(`{{define "T"}}http/template: {{.}}!{{end}}`) // what is it?
		t1.ExecuteTemplate(w, "T", template.HTML(input))

		//output converted js to client
		//template.HTMLEscape(w, []byte(input))
	}
}

