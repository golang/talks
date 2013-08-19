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

// Example of bad code, missing early return. OMIT
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err == nil {
		size += 4
		var n int
		n, err = w.Write([]byte(g.Name))
		size += int64(n)
		if err == nil {
			err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
			if err == nil {
				size += 4
			}
			return
		}
		return
	}
	return
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

// Example of bad API, it's better to use an interface.
func (g *Gopher) WriteToFile(f *os.File) (int64, error) {
	return 0, nil
}

// Example of bad API, it's better to use a narrower interface.
func (g *Gopher) WriteToReadWriter(rw io.ReadWriter) (int64, error) {
	return 0, nil
}

// Example of better API.
func (g *Gopher) WriteToWriter(f io.Writer) (int64, error) {
	return 0, nil
}
