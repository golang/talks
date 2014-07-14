// +build ignore

package main

import "fmt"

// 1 START OMIT
const C1 = 1e-323

const C2 = C1 / 100
const C3 = C2 * 100

const C4 float64 = C1 / 100
const C5 = C4 * 100

func main() {
	fmt.Println(C3, C5)
}

// 1 END OMIT
