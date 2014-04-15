// +build OMIT

package main

import "fmt"

func main() {
	c := make(chan string)
	go func() {
		c <- "Hello"
		c <- "World"
	}()
	fmt.Println(<-c, <-c)
}
