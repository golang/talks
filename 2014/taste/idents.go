// +build OMIT

package main // idents.go

import (
	"fmt"
	"os"
	"text/scanner"
)

func main() {
	var s scanner.Scanner
	s.Init(os.Stdin)
	for {
		switch s.Scan() {
		case scanner.EOF:
			return // all done
		case scanner.Ident:
			fmt.Println(s.TokenText())
		}
	}
}
