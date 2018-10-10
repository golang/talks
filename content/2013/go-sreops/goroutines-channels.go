// +build OMIT

package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	textChannel := make(chan string)
	words := []string{"ho!", "hey!"}
	secs := []int{2, 1}
	// Create a goroutine per word
	for i, word := range words {
		go say(word, secs[i], textChannel) // &
	}
	// Wait for response via channel N times
	for _ = range words {
		fmt.Println(<-textChannel)
	}
}

// say sends word back via channel after sleeping for X secs
func say(word string, secs int, textChannel chan string) {
	time.Sleep(time.Duration(secs) * time.Second)
	textChannel <- word
}

// STOP OMIT
