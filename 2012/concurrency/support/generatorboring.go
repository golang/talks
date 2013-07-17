// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
// START1 OMIT
	c := boring("boring!") // Function returning a channel. // HL
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring; I'm leaving.")
// STOP1 OMIT
}

// START2 OMIT
func boring(msg string) <-chan string { // Returns receive-only channel of strings. // HL
	c := make(chan string)
	go func() { // We launch the goroutine from inside the function. // HL
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // Return the channel to the caller. // HL
}
// STOP2 OMIT

