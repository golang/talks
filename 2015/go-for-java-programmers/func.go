// +build OMIT

package main

import "fmt"

// div divides n by d and returns the quotient and remainder.
func div(n, d int) (q, r int) { // HL
	return n / d, n % d
}

func main() {
	quot, rem := div(4, 3)
	fmt.Println(quot, rem)
}
