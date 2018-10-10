// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func main() {
	c := boring("Joe")
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-time.After(1 * time.Second): // HL
			fmt.Println("You're too slow.")
			return
		}
	}
}
// STOP1 OMIT

// START2 OMIT
func boring(msg string) <-chan string { // Returns receive-only channel of strings. // HL
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller. // HL
}
// STOP2 OMIT


// START3 OMIT
func fanIn(input1, input2 <-chan string) <-chan string { // HL
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
// STOP3 OMIT
