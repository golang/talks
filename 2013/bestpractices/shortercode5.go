// +build ignore,OMIT

package main

import (
	"encoding/binary"
	"image/color"
	"io"
	"log"
	"os"
)

// GOPHER OMIT
type Gopher struct {
	Name     string
	Age      int32
	FurColor color.Color
}

type binWriter struct {
	w   io.Writer
	err error
}

// WRITE OMIT
// Write write the given value into the writer using little endian.
func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	switch v := v.(type) {
	case string:
		w.Write(int32(len(v)))
		w.Write([]byte(v))
	default:
		w.err = binary.Write(w.w, binary.LittleEndian, v)
	}
}

// DUMP OMIT
func (g *Gopher) DumpBinary(w io.Writer) error {
	bw := &binWriter{w: w}
	bw.Write(g.Name)
	bw.Write(g.Age)
	bw.Write(g.FurColor)
	return bw.err
}

// MAIN OMIT
func main() {
	w := os.Stdout
	g := &Gopher{
		Name:     "Gophertiti",
		Age:      3383,
		FurColor: color.RGBA{B: 255},
	}

	if err := g.DumpBinary(w); err != nil {
		log.Fatal("DumpBinary: %v", err)
	}
}
