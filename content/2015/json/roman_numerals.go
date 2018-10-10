// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type romanNumeral int

var numerals = []struct {
	s string
	v int
}{
	{"M", 1000}, {"CM", 900},
	{"D", 500}, {"CD", 400},
	{"C", 100}, {"XC", 90},
	{"L", 50}, {"XL", 40},
	{"X", 10}, {"IX", 9},
	{"V", 5}, {"IV", 4},
	{"I", 1},
}

func (n romanNumeral) String() string {
	res := ""
	v := int(n)
	for _, num := range numerals {
		res += strings.Repeat(num.s, v/num.v)
		v %= num.v
	}
	return res
}

func parseRomanNumeral(s string) (romanNumeral, error) {
	res := 0
	for _, num := range numerals {
		for strings.HasPrefix(s, num.s) {
			res += num.v
			s = s[len(num.s):]
		}
	}
	return romanNumeral(res), nil
}

func (n romanNumeral) MarshalJSON() ([]byte, error) {
	if n <= 0 {
		return nil, fmt.Errorf("Romans had only natural (=>1) numbers")
	}
	return json.Marshal(n.String())
}

func (n *romanNumeral) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	p, err := parseRomanNumeral(s)
	if err == nil {
		*n = p
	}
	return err
}

type Movie struct {
	Title string
	Year  romanNumeral
}

func main() {
	// Encoding
	movies := []Movie{{"E.T.", 1982}, {"The Matrix", 1999}, {"Casablanca", 1942}}
	res, err := json.MarshalIndent(movies, "", "\t") // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Movies: %s\n", res)

	// Decoding
	var m Movie
	inputText := `{"Title": "Alien", "Year":"MCMLXXIX"}`
	if err := json.NewDecoder(strings.NewReader(inputText)).Decode(&m); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s was released in %d\n", m.Title, m.Year)
}
