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

const input = `{
    "name":"Gopher",
    "birthdate": "2009/11/10",
    "shirt-size": "XS"
}`

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

func (ss *ShirtSize) UnmarshalJSON(data []byte) error {
	// Extract the string from data.
	var s string
	if err := json.Unmarshal(data, &s); err != nil { // HL
		return fmt.Errorf("shirt-size should be a string, got %s", data)
	}

	// The rest is equivalen to ParseShirtSize.
	got, ok := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}[s]
	if !ok {
		return fmt.Errorf("invalid ShirtSize %q", s)
	}
	*ss = got // HL
	return nil
}

func (p *Person) UnmarshalJSON(data []byte) error {
	var aux struct {
		Name string
		Born string    `json:"birthdate"`
		Size ShirtSize `json:"shirt-size"` // HL
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	if err := dec.Decode(&aux); err != nil {
		return fmt.Errorf("decode person: %v", err)
	}
	p.Name = aux.Name
	p.Size = aux.Size // HL
	// ... rest of function omitted ...
	born, err := time.Parse("2006/01/02", aux.Born)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	p.Born = born
	return nil
}

func main() {
	var p Person
	dec := json.NewDecoder(strings.NewReader(input))
	if err := dec.Decode(&p); err != nil {
		log.Fatalf("parse person: %v", err)
	}
	fmt.Println(p)
}
