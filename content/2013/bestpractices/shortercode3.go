// +build OMIT

package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

// Write writes a value to the provided writer in little endian form.
func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
		w.size += int64(binary.Size(v))
	}
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(int32(len(g.Name)))
	bw.Write([]byte(g.Name))
	bw.Write(int64(g.AgeYears))
	return bw.size, bw.err
}

func main() {
	g := &Gopher{
		Name:     "Gophertiti",
		AgeYears: 3382,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}
