// +build OMIT

package main

import (
	"fmt"
	"time"
)

func sleepAndTalk(t time.Duration, msg string) {
	time.Sleep(t)
	fmt.Printf("%v ", msg)
}

func main() {
	go sleepAndTalk(0*time.Second, "Hello")
	go sleepAndTalk(1*time.Second, "Gophers!")
	go sleepAndTalk(2*time.Second, "What's")
	go sleepAndTalk(3*time.Second, "up?")
	time.Sleep(4 * time.Second)
}
