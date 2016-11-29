package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	// for prevent repeat create token
	"crypto/md5"
	"io"
	"time"
)

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
