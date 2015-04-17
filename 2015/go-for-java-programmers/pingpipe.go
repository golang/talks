// +build OMIT

package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{ hits int }

func main() {
	in, out := make(chan *Ball), make(chan *Ball) // HL
	go player("ping", in, out)
	go player("pong", in, out)

	go func() {
		for i := 0; i < 8; i++ {
			in <- new(Ball) // feed the pipeline // HL
		}
	}()
	for i := 0; i < 8; i++ {
		<-out // drain the pipeline // HL
	}
}

func player(name string, in <-chan *Ball, out chan<- *Ball) { // HL
	for i := 0; ; i++ {
		ball := <-in // HL
		ball.hits++
		fmt.Println(name, i, "hit", ball.hits)
		time.Sleep(100 * time.Millisecond)
		out <- ball // HL
	}
}

// STOPMAIN1 OMIT
