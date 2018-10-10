// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(10)
}

func sendMessages() chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			ch <- fmt.Sprintf("message %v", i)
		}
	}()
	return ch
}

func main() {
	timeout := time.NewTimer(80 * time.Millisecond)
	ch := sendMessages()
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			timeout.Reset(80 * time.Millisecond)
		case <-timeout.C:
			fmt.Println("timeout")
			return
		}
	}
}
