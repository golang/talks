// +build ignore

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var s interface {
		String() string
	} = bytes.NewBufferString("I'm secretly a fmt.Stringer!")
	fmt.Println(s.String())
}
