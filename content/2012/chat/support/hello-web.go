// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
)

const listenAddr = "localhost:4000"

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web")
}
