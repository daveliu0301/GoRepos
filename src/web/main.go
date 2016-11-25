package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	textTemplate "text/template"

	// for prevent repeat create token
	"crypto/md5"
	"io"
	"time"

	//upload
	"os"
)

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("handleUpload.gtpl")
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
		f, err := os.OpenFile("./test/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)

	}
}

func getFileName(p string) string {
	arr :=	strings.Split(p, ".")
	if len(arr) >= 1 {
		return getToken() + "." + arr[len(arr) - 1]
	} else {
		return getToken()
	}
}

func getToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	return fmt.Sprintf("%x", h.Sum(nil))

}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello handsome guy!")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

func checkInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("checkInput.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()

		required := r.Form["required"][0]
		if isEmpty(required) {
			fmt.Println("please input required")
		} else {
			fmt.Println("required: ", required)
		}

		number := r.Form["number"][0]
		if isEmpty(number) {
			fmt.Println("number is empty")
		} else if isInteger(number) {
			i, _ := strconv.Atoi(number)
			fmt.Println("number: ", i)
		} else {
			fmt.Println("please input 0-9")
		}

		chinese := r.Form["chinese"][0]
		if isEmpty(chinese) {
			fmt.Println("chinese is empty")
		} else if isChinese(chinese) {
			fmt.Println("chinese: ", chinese)
		} else {
			fmt.Println("please input chinese")
		}

		english := r.Form["english"][0]
		if isEmpty(english) {
			fmt.Println("english is empty")
		} else if isEnglish(english) {
			fmt.Println("english: ", english)
		} else {
			fmt.Println("please input english")
		}

		fruit := r.Form.Get("fruit")
		if checkSelect(fruit) {
			fmt.Println("fruit: ", fruit)
		} else {
			fmt.Println("wrong input: ", fruit)
		}

		gender := r.Form.Get("gender")
		if checkGender(gender) {
			fmt.Println("gender: ", gender)
		} else {
			fmt.Println("wrong input: ", gender)
		}

		interest := r.Form["interest"]
		m, s := checkCheckBox(interest)
		if m {
			fmt.Println("interest: ", interest)
		} else {
			fmt.Println("wrong interest input: ", s)
		}
	}
}

func preventRepeat(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("preventRepeat.gtpl")
		t.Execute(w, getToken()) // second arg?
	} else {
		r.ParseForm()
		fmt.Fprintln(w, r.Form["input"][0])
		fmt.Fprintln(w, r.Form.Get("token"))
		fmt.Println(r.Form["input"][0])
		fmt.Println(r.Form.Get("token")) // different and usage r.Form.Get() & r.Form[][]
	}

}

func preventXSS(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("preventXSS.gtpl")
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

//check checkBox input
func checkCheckBox(slice []string) (bool, string) {
	for _, v := range slice {
		m, _ := regexp.MatchString("^(football|basketball)$", v)
		if !m {
			return false, v
		}
	}
	return true, ""
}

//check radio box
func checkGender(s string) bool {
	m, _ := regexp.MatchString("^(1|2)$", s)
	return m
}

//check option menu
func checkSelect(s string) bool {
	m, _ := regexp.MatchString("^(apple|pear|banana)$", s)
	return m
}

func isEmpty(s string) bool {
	return len(s) == 0
}

func isInteger(s string) bool {
	m, _ := regexp.MatchString("^[0-9]+$", s)
	return m
}

func isChinese(s string) bool {
	m, _ := regexp.MatchString("^\\p{Han}+$", s)
	return m
}

func isEnglish(s string) bool {
	m, _ := regexp.MatchString("^[a-zA-Z]+$", s)
	return m
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/checkInput", checkInput)
	http.HandleFunc("/preventRepeat", preventRepeat)
	http.HandleFunc("/preventXSS", preventXSS)
	http.HandleFunc("/handleUpload", handleUpload)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
