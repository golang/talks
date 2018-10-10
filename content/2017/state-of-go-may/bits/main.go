// +build go1.9

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	const n = 100
	// START OMIT
	fmt.Printf("%d (%b) has %d bits set to one\n", n, n, bits.OnesCount(n))

	fmt.Printf("%d reversed is %d\n", n, bits.Reverse(n))

	fmt.Printf("%d can be encoded in %d bits\n", n, bits.Len(n))
	// END OMIT
}
