// +build OMIT

package main

import (
	"errors"
	"fmt"
)

// div divides n by d and returns the quotient and remainder.
// It returns an error if d is zero.
func div(n, d int) (q, r int, err error) { // HL
	if d == 0 {
		err = errors.New("divide by zero") // HL
		return
	}
	return n / d, n % d, nil // HL
}

func main() {
	fmt.Println(div(4, 3))
	fmt.Println(div(3, 0))
}
