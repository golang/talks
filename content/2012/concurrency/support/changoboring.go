// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
func main() {
	c := make(chan string)
	go boring("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value. // HL
	}
	fmt.Println("You're boring; I'm leaving.")
}
// STOP1 OMIT

// START2 OMIT
func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value. // HL
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}
// STOP2 OMIT

