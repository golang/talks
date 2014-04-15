// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	go say("ho!", 2*time.Second)  // &
	go say("hey!", 1*time.Second) // &

	// Make main sleep for 4 seconds so goroutines can finish
	time.Sleep(4 * time.Second)
}

// say prints text after sleeping for X secs
func say(text string, secs time.Duration) {
	time.Sleep(secs)
	fmt.Println(text)
}
