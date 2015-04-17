// +build OMIT

package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{ hits int }

func main() {
	in, out := make(chan *Ball), make(chan *Ball)
	go player("ping", in, out)
	go player("pong", in, out)

	for i := 0; i < 8; {
		select { // HL
		case in <- new(Ball): // feed the pipeline // HL
		case <-out: // drain the pipeline // HL
			i++ // HL
		} // HL
	}
}

func player(name string, in <-chan *Ball, out chan<- *Ball) {
	for i := 0; ; i++ {
		ball := <-in
		ball.hits++
		fmt.Println(name, i, "hit", ball.hits)
		time.Sleep(100 * time.Millisecond)
		out <- ball
	}
}

// STOPMAIN1 OMIT
