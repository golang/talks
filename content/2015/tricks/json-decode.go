// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data struct {
		ID   int
		Name string
	}
	err := json.Unmarshal([]byte(`{"ID": 42, "Name": "The answer"}`), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.ID, data.Name)
}
