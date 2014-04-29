// +build OMIT

package main

import (
	"fmt"
	"io"
	"log"
)

// ByteReader implements an io.Reader that emits a stream of its byte value.
type ByteReader byte

func (b ByteReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = byte(b)
	}
	return len(buf), nil
}

type LogReader struct {
	io.Reader
}

func (r LogReader) Read(b []byte) (int, error) {
	n, err := r.Reader.Read(b)
	log.Printf("read %d bytes, error: %v", n, err)
	return n, err
}

// STOP OMIT

func main() {
	// START OMIT
	r := LogReader{ByteReader('A')}
	b := make([]byte, 10)
	r.Read(b)
	fmt.Printf("b: %q", b)
	// STOP OMIT
}
