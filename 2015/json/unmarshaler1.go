// +build OMIT

package main

import (
	"bytes"
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
	s, ok := map[ShirtSize]string{XS: "XS", S: "S", M: "M", L: "L", XL: "XL"}[ss]
	if !ok {
		return "invalid ShirtSize"
	}
	return s
}

func ParseShirtSize(s string) (ShirtSize, error) {
	ss, ok := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}[s]
	if !ok {
		return NA, fmt.Errorf("invalid ShirtSize %q", s)
	}
	return ss, nil
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name string
		Born string `json:"birthdate"`
		Size string `json:"shirt-size"`
	}

	dec := json.NewDecoder(bytes.NewReader(data)) // HL
	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}
	p.Name = aux.Name
	// ... rest of function omitted ...
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
	dec := json.NewDecoder(strings.NewReader(input))
	if err := dec.Decode(&p); err != nil {
		log.Fatalf("parse person: %v", err)
	}
	fmt.Println(p)
}
