// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START1 OMIT
func main() {
	c := make(chan string)
	go f("three", 300*time.Millisecond, c) // HL
	go f("six", 600*time.Millisecond, c)   // HL
	go f("nine", 900*time.Millisecond, c)  // HL
	for i := 0; i < 10; i++ {
		fmt.Println("Received", <-c)
	}
	fmt.Println("Done.")
}

// STOP1 OMIT

// START2 OMIT
func f(msg string, delay time.Duration, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Any suitable value can be sent. // HL
		time.Sleep(delay)
	}
}

// STOP2 OMIT
