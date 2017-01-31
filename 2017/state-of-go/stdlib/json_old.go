package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	in := []byte(`
    {
        "full_name": "Gopher",
        "age": 7,
        "social_security": 1234
    }`)

	var p Person
	if err := json.Unmarshal(in, &p); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", p)
}

type Person struct {
	Name     string
	AgeYears int
	SSN      int
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name     string `json:"full_name"`
		AgeYears int    `json:"age"`
		SSN      int    `json:"social_security"`
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*p = Person{
		Name:     aux.Name,
		AgeYears: aux.AgeYears,
		SSN:      aux.SSN,
	}
	*p = Person(aux)
	return nil
}
