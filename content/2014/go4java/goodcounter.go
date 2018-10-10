// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

var nextID = make(chan int)

func handler(w http.ResponseWriter, q *http.Request) {
	fmt.Fprintf(w, "<h1>You got %v<h1>", <-nextID)
}

func main() {
	http.HandleFunc("/next", handler)
	go func() {
		for i := 0; ; i++ {
			nextID <- i
		}
	}()
	http.ListenAndServe("localhost:8080", nil)
}
