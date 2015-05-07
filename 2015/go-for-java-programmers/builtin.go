// +build OMIT

package main

import (
	"fmt"
	"sort"
)

func main() {
	hellos := map[string]string{ // HLbuiltin
		"English":  "Hello",   // HLstrings
		"Mandarin": "您好",      // HLstrings
		"Hindi":    "नमस्कार", // HLstrings
	}
	fmt.Println(hellos)               // HLfmt
	var langs []string                // HLbuiltin
	for lang, hello := range hellos { // HLbuiltin
		fmt.Println(lang, ":", hello, "world!") // HLfmt
		langs = append(langs, lang)             // HLbuiltin
	}
	sort.Strings(langs)                           // HLstrings
	fmt.Printf("len(%v) = %d", langs, len(langs)) // HLfmt
}
