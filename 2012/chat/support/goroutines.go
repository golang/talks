// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	go say("let's go!", 3)
	go say("ho!", 2)
	go say("hey!", 1)
	time.Sleep(4 * time.Second)
}

func say(text string, secs int) {
	time.Sleep(time.Duration(secs) * time.Second)
	fmt.Println(text)
}
