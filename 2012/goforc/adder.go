// +build OMIT

package main

import "fmt"

// START1 OMIT
func adder(delta int) func(x int) int {
	f := func(x int) int { // HL
		return x + delta // HL
	} // HL
	return f
}

// STOP1 OMIT

func main() {
	// START2 OMIT
	var inc = adder(1)
	fmt.Println(inc(0))
	fmt.Println(adder(-1)(10))
	// STOP2 OMIT
}
