// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

var authURL = ""

var auth = func(user string) bool {
	res, err := http.Get(authURL + "/" + user)
	return err == nil && res.StatusCode == http.StatusOK
}

func sayHi(user string) {
	if !auth(user) {
		fmt.Printf("unknown user %v\n", user)
		return
	}
	fmt.Printf("Hi, %v\n", user)
}

func TestSayHi() {
	auth = func(string) bool { return true }
	sayHi("John")

	auth = func(string) bool { return false }
	sayHi("John")
}

func init() {
	auth = func(string) bool { return true }
}

func TestAnythingElse() {
	// auth has been already set to the fake version
}

func main() {
	TestSayHi()
}
