// +build ignore

package main

import "fmt"

// 2 START OMIT
// 1 START OMIT
func main() {
	i := 1
	f := func() T {
		return T{
			i: 1, // HL
		}
	}
	fmt.Println(i, f())
}

// 1 END OMIT

type T struct{ i int }

// 2 END OMIT
