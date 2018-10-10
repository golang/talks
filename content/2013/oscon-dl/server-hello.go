// +build ignore,OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stdout, "%s details: %+v\n", r.URL.Path, r)
	fmt.Fprintf(w, "Hello, world! at %s\n", r.URL.Path)
}

func main() {
	log.Printf("Running...")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", http.HandlerFunc(handler)))
}
