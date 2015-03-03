// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

const input = `
    {
        "name":"Gopher",
        "birthdate": "2009/11/10",
        "shirt-size": "XS"
    }
    `

type Person struct {
	Name string    `json:"name"`
	Born time.Time `json:"birthdate"`
	Size ShirtSize `json:"shirt-size"`
}

type ShirtSize byte

const (
	NA ShirtSize = iota
	XS
	S
	M
	L
	XL
)

func (ss ShirtSize) String() string {
	s, ok := map[ShirtSize]string{XS: "XS", S: "S", M: "M", L: "L", XL: "XL"}[ss]
	if !ok {
		return "invalid ShirtSize"
	}
	return s
}

func main() {
	var p Person
	dec := json.NewDecoder(strings.NewReader(input))
	if err := dec.Decode(&p); err != nil {
		log.Fatalf("parse person: %v", err)
	}
	fmt.Println(p)
}
