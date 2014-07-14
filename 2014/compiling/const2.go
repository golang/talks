// +build ignore

package main

import "fmt"

// 1 START OMIT
const C1 = 1e+308
const C2 = C1 * 10
const C3 = C2 / 10

var V1 = C1
var V2 = V1 * 10
var V3 = V2 / 10

func main() {
	fmt.Println(C3, V3)
}

// 1 END OMIT
