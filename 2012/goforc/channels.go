// +build OMIT

package main

func main() {
	var value int

	// START1 OMIT
	// Declaring and initializing.
	var c chan int
	c = make(chan int)
	// or
	c := make(chan int) // HL
	// STOP1 OMIT

	// START2 OMIT
	// Sending on a channel.
	c <- 1 // HL
	// STOP2 OMIT

	// START3 OMIT
	// Receiving from a channel.
	// The "arrow" indicates the direction of data flow.
	value = <-c // HL
	// STOP3 OMIT

	_ = value
}
