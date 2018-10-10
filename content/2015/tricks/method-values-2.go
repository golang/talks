// +build ignore

package main

import (
	"bytes"
	"os"
)

func main() {
	var f func(string) (int, error)
	var buf bytes.Buffer
	f = buf.WriteString
	f("Hey... ")
	f("this *is* cute.")
	buf.WriteTo(os.Stdout)
}
