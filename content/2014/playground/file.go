// +build OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	const filename = "/tmp/file.txt"

	err := ioutil.WriteFile(filename, []byte("Hello, file system\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
