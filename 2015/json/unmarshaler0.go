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
	Name string
	Born time.Time
	Size ShirtSize
}

func (p Person) String() string {
	return fmt.Sprintf("%s was born on %v and uses a %v t-shirt",
		p.Name, p.Born.Format("Jan 2 2006"), p.Size)
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
	sizes := map[ShirtSize]string{XS: "XS", S: "S", M: "M", L: "L", XL: "XL"}
	s, ok := sizes[ss]
	if !ok {
		return "invalid t-shirt size"
	}
	return s
}

func ParseShirtSize(s string) (ShirtSize, error) {
	sizes := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}
	ss, ok := sizes[s]
	if !ok {
		return NA, fmt.Errorf("invalid t-shirt size %q", s)
	}
	return ss, nil
}

func (p *Person) Parse(s string) error {
	var aux struct {
		Name string
		Born string `json:"birthdate"`
		Size string `json:"shirt-size"`
	}

	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	p.Name = aux.Name
	born, err := time.Parse("2006/01/02", aux.Born)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	p.Born = born
	p.Size, err = ParseShirtSize(aux.Size)
	return err
}

func main() {
	var p Person
	if err := p.Parse(input); err != nil {
		log.Fatalf("parse person: %v", err)
	}
	fmt.Println(p)
}
