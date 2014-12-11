package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)
	c1 <- 42

	select {
	case v := <-c1:
		fmt.Println("received from c1: ", v)
	case c2 <- 1:
		fmt.Println("sent to c2")
	case <-time.After(time.Second):
		fmt.Println("timed out")
	default:
		fmt.Println("nothing ready at the moment")
	}
}
