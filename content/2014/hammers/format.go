// +build ignore

package main

import (
	"fmt"
	"go/format"
)

func main() {
	ugly := `func (f *File) Read(p []byte, )(n int, err error, ){}`
	fmt.Println(ugly)
	pretty, _ := format.Source([]byte(ugly)) // HL
	fmt.Println(string(pretty))
}
