package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	resp, err := http.Get("https://bing.com")
	if err != nil{
		fmt.Println(err)
		return	
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		fmt.Println(err)
		return
	}

	fmt.Printf("status: %d\n", resp.StatusCode)

	err = ioutil.WriteFile("./temp/data", data, 0664)
	if err != nil{
		fmt.Println(err)
		return
	}
}
