package main

import (
	"flag"
	"fmt"
)

var methods = map[string]func(){
	"login":  login,
	"auth":   auth,
	"buy":    buy,
	"logout": logout,
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("input one arg at least.")
		return
	}

	for i := 0; i < flag.NArg(); i++ {
		handler := methods[flag.Arg(i)]
		if handler != nil {
			handler()
			return
		}
	}
	fmt.Println("finish.")
}

func login() {
	fmt.Println("handle login")
}

func auth() {
	fmt.Println("handle auth")
}

func buy() {
	fmt.Println("handle buy")
}

func logout() {
	fmt.Println("handle logout")
}
