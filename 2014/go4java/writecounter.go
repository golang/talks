// +build OMIT

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

var (
	_ = bytes.Buffer{}
	_ = os.Stdout
)

// WriteCounter tracks the total number of bytes written.
type WriteCounter struct {
	io.ReadWriter
	count int
}

func (w *WriteCounter) Write(b []byte) (int, error) {
	w.count += len(b)
	return w.ReadWriter.Write(b)
}

// MAIN OMIT
func main() {
	buf := &bytes.Buffer{}
	w := &WriteCounter{ReadWriter: buf}

	fmt.Fprintf(w, "Hello, gophers!\n")
	fmt.Printf("Printed %v bytes", w.count)
}
