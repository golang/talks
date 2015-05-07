// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // HL
	fmt.Fprintln(w, "hello")
}

func main() {
	http.HandleFunc("/", handler) // HL
	err := http.ListenAndServe("localhost:1234", nil)
	if err != nil {
		log.Fatal(err)
	}
}
