// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	var data struct {
		ID     int
		Person struct {
			Name string
			Job  string
		}
	}
	const s = `{"ID":42,"Person":{"Name":"George Costanza","Job":"Architect"}}`
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(data.ID, data.Person.Name, data.Person.Job)
}
