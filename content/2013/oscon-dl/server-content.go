// +build ignore,OMIT

package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	log.Printf("Running...")
	err := http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "foo.txt", time.Now(),
			strings.NewReader("I am some content.\n"))
	}))
	log.Fatal(err)
}
