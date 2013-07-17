// +build OMIT

package main

import (
	"io"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func slurp(r io.Reader) error {
	b := make([]byte, 1024)
	for {
		_, err := r.Read(b)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	panic("unreachable")
}

func main() {
	println(min(10, 20))
	slurp(os.Stdin)
}
