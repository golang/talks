// +build ignore,OMIT

package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.Printf("Running...")
	log.Fatal(http.ListenAndServe(
		"127.0.0.1:8080",
		http.FileServer(http.Dir(
			filepath.Join(os.Getenv("HOME"), "go", "doc")))))
}
