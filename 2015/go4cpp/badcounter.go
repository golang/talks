// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID int

func handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v<h1>", nextID)
	nextID++
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
