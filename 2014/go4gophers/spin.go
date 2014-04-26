package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

type Zero struct{}

func (Zero) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 0
	}
	return len(b), nil
}

type SleepReader struct {
	r io.Reader
	d time.Duration
}

func (r SleepReader) Read(b []byte) (int, error) {
	time.Sleep(r.d)
	return r.r.Read(b)
}

// END OMIT

type Spinner int

var spinBytes = []byte{'-', '/', '|', '\\'}

func (s *Spinner) Tick() {
	*s = (*s + 1) % 4
	fmt.Printf("\x0cReading...%c", spinBytes[*s])
}

type SpinReader struct {
	r io.Reader
	s Spinner
}

func (r *SpinReader) Read(b []byte) (int, error) {
	r.s.Tick()
	return r.r.Read(b)
}

func main() {
	var r io.Reader = Zero{}
	r = io.LimitReader(r, 1e6)
	r = SleepReader{r, 20 * time.Millisecond}
	r = &SpinReader{r: r}
	n, _ := io.Copy(ioutil.Discard, r)
	fmt.Print("\x0c", n, " bytes read.")
}
