// +build OMIT

package main

import (
	"fmt"
	"github.com/nf/reddit" // HL
	"log"
)

func main() {
	items, err := reddit.Get("golang") // HL
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
