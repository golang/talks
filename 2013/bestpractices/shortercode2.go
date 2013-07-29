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

// DUMP OMIT
func (g *Gopher) DumpBinary(w io.Writer) error {
	err := binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(g.Name))
	if err != nil {
		return err
	}
	err = binary.Write(w, binary.LittleEndian, g.Age)
	if err != nil {
		return err
	}
	return binary.Write(w, binary.LittleEndian, g.FurColor)
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
