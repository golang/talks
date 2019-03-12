// +build ignore,OMIT

package main

func main() {
	c := make(chan int)

	<-c
}
