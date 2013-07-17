// +build OMIT

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const blob = `[
	{"Title":"Øredev", "URL":"http://oredev.org"},
	{"Title":"Strange Loop", "URL":"http://thestrangeloop.com"}
]`

type Item struct {
	Title string
	URL   string
}

func main() {
	var items []*Item
	json.NewDecoder(strings.NewReader(blob)).Decode(&items)
	for _, item := range items {
		fmt.Printf("Title: %v URL: %v\n", item.Title, item.URL)
	}
}
