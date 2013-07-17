// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func cleanup() {
}

func main() {
// START1 OMIT
	quit := make(chan string) // HL
	c := boring("Joe", quit) // HL
	for i := rand.Intn(10); i >= 0; i-- { fmt.Println(<-c) }
	quit <- "Bye!" // HL
	fmt.Printf("Joe says: %q\n", <-quit) // HL
// STOP1 OMIT
}

func boring(msg string, quit chan string) <-chan string {
	c := make(chan string) // HL
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
// START2 OMIT
			select {
			case c <- fmt.Sprintf("%s: %d", msg, i):
				// do nothing
			case <-quit: // HL
				cleanup()
				quit <- "See you!" // HL
				return
			}
// STOP2 OMIT
		}
	}()
	return c
}
