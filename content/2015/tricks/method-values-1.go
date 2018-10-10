// +build ignore

package main

import (
	"bytes"
	"os"
)

func main() {
	var f func(*bytes.Buffer, string) (int, error)
	var buf bytes.Buffer
	f = (*bytes.Buffer).WriteString
	f(&buf, "y u no buf.WriteString?")
	buf.WriteTo(os.Stdout)
}
