// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v<h1>", <-nextID)
}

func counter() {
	for i := 0; ; i++ {
		nextID <- i
	}
}

func main() {
	http.HandleFunc("/", handler)
	go counter()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
