// +build OMIT

package main

import "fmt"

var primes = [...]int{2, 3, 5, 7, 11, 13, 17, 19}

func _() {
	// START1 OMIT
	for i := 0; i < len(primes); i++ {
		fmt.Println(i, primes[i])
	}
	// STOP1 OMIT

	// START2 OMIT
	var sum int
	for _, x := range primes {
		sum += x
	}
	// STOP2 OMIT
}

func main() {
	// START3 OMIT
	for i, x := range primes {
		fmt.Println(i, x)
	}
	// STOP3 OMIT
}
