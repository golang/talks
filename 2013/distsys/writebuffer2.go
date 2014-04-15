// +build OMIT

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "hello, %s\n", "world")
	io.Copy(os.Stdout, b)
}
