package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", sayHelloName)
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
