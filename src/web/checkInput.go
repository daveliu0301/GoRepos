package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func checkInput(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./gtpls/checkInput.gtpl")
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
