// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	go f("three", 300*time.Millisecond)
	go f("six", 600*time.Millisecond)
	go f("nine", 900*time.Millisecond)
}

// STOP OMIT

func f(msg string, delay time.Duration) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(delay)
	}
}
