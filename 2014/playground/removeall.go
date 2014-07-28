// +build OMIT

package main

import (
	"log"
	"os"
)

func main() {
	err := os.RemoveAll("/foo")
	if err != nil {
		log.Fatal(err)
	}
}
