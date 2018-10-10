// +build OMIT

package main

import (
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
	Name string    `json:"name"`
	Born Date      `json:"birthdate"`
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

func (ss *ShirtSize) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("shirt-size should be a string, got %s", data)
	}
	got, ok := map[string]ShirtSize{"XS": XS, "S": S, "M": M, "L": L, "XL": XL}[s]
	if !ok {
		return fmt.Errorf("invalid ShirtSize %q", s)
	}
	*ss = got
	return nil
}

type Date struct{ time.Time }

func (d Date) String() string { return d.Format("2006/01/02") }

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("birthdate should be a string, got %s", data)
	}
	t, err := time.Parse("2006/01/02", s)
	if err != nil {
		return fmt.Errorf("invalid date: %v", err)
	}
	d.Time = t
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
