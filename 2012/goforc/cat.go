// +build OMIT

package main

import (
	"flag"
	"io"
	"os"
)

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		f, err := os.Open(arg)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		_, err = io.Copy(os.Stdout, f) // HL
		if err != nil {
			panic(err)
		}
	}
}
