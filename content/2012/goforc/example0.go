// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START1 OMIT
func main() {
	f("Hello, World", 500*time.Millisecond)
}

// STOP1 OMIT

// START2 OMIT
func f(msg string, delay time.Duration) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(delay)
	}
}

// STOP2 OMIT
