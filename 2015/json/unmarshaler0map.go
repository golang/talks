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
    "name": "Gopher",
    "birthdate": "2009/11/10",
    "shirt-size": "XS"
}
`

type Person struct {
	Name string
	Born time.Time
	Size ShirtSize
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
		return "invalid ShirtSize"
	}
	return s
}

func ParseShirtSize(s string) (ShirtSize, error) {
	sizes := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}
	ss, ok := sizes[s]
	if !ok {
		return NA, fmt.Errorf("invalid ShirtSize %q", s)
	}
	return ss, nil
}

func (p *Person) Parse(s string) error {
	fields := map[string]string{}

	dec := json.NewDecoder(strings.NewReader(s))
	if err := dec.Decode(&fields); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}

	// Once decoded we can access the fields by name.
	p.Name = fields["name"]

	born, err := time.Parse("2006/01/02", fields["birthdate"])
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	p.Born = born

	p.Size, err = ParseShirtSize(fields["shirt-size"])
	return err
}

func main() {
	var p Person
	if err := p.Parse(input); err != nil {
		log.Fatalf("parse person: %v", err)
	}
	fmt.Println(p)
}
