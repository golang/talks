package main

import (
	"io"
	"io/ioutil"
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

func main() {
	r := LogReader{ByteReader('A')}
	io.CopyN(ioutil.Discard, r, 1e6)
}
