// +build OMIT

package main

import (
	"fmt"
)

func main() {
	in, out := make(chan int), make(chan int)
	go buffer(in, out)
	for i := 0; i < 10; i++ {
		in <- i
	}
	close(in)
	for i := range out {
		fmt.Println(i)
	}
}

// buffer provides an unbounded buffer between in and out.  buffer
// exits when in is closed and all items in the buffer have been sent
// to out, at which point it closes out.
func buffer(in <-chan int, out chan<- int) {
	var buf []int
	for in != nil || len(buf) > 0 {
		var i int
		var c chan<- int
		if len(buf) > 0 {
			i = buf[0]
			c = out // enable send case
		}
		select {
		case n, ok := <-in:
			if ok {
				buf = append(buf, n)
			} else {
				in = nil // disable receive case
			}
		case c <- i:
			buf = buf[1:]
		}
	}
	close(out)
}
