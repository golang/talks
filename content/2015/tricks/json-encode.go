// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	b, err := json.Marshal(struct {
		ID   int
		Name string
	}{42, "The answer"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
