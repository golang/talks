// +build OMIT

package main

import (
	"fmt"
	"time"
)

// f START OMIT
func f(msg string, delay time.Duration, ch chan string) {
	for {
		ch <- msg
		time.Sleep(delay)
	}
}

// f END OMIT

// main START OMIT
func main() {
	ch := make(chan string)
	go f("A--", 300*time.Millisecond, ch)
	go f("-B-", 500*time.Millisecond, ch)
	go f("--C", 1100*time.Millisecond, ch)

	for i := 0; i < 100; i++ {
		fmt.Println(i, <-ch)
	}
}

// main END OMIT
