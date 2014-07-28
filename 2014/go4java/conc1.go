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
	sleepAndTalk(0*time.Second, "Hello")
	sleepAndTalk(1*time.Second, "Gophers!")
	sleepAndTalk(2*time.Second, "What's")
	sleepAndTalk(3*time.Second, "up?")
}
