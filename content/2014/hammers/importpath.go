// +build ignore

package main

import (
	"fmt"

	"golang.org/x/tools/imports"
)

func main() {
	iface := "http.Handler"
	src := "package hack; var i " + iface // HL
	fmt.Println(src, "\n---")

	imp, _ := imports.Process("", []byte(src), nil) // HL
	// ignoring errors throughout this presentation
	fmt.Println(string(imp))
}
