// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START0 OMIT
type Message struct {
	str string
	wait chan bool // HL
}
// STOP0 OMIT

func main() {
	c := fanIn(boring("Joe"), boring("Ann")) // HL
// START1 OMIT
	for i := 0; i < 5; i++ {
		msg1 := <-c; fmt.Println(msg1.str)
		msg2 := <-c; fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
// STOP1 OMIT
	fmt.Println("You're all boring; I'm leaving.")
}

func boring(msg string) <-chan Message { // Returns receive-only channel of strings. // HL
	c := make(chan Message)
// START2 OMIT
	waitForIt := make(chan bool) // Shared between all messages.
// STOP2 OMIT
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
// START3 OMIT
			c <- Message{ fmt.Sprintf("%s: %d", msg, i), waitForIt }
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
// STOP3 OMIT
		}
	}()
	return c // Return the channel to the caller. // HL
}


// START3 OMIT
func fanIn(inputs ... <-chan Message) <-chan Message { // HL
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i] // New instance of 'input' for each loop.
		go func() { for { c <- <-input } }()
	}
	return c
}
// STOP3 OMIT
