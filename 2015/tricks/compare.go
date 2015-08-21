// +build ignore

package main

import "fmt"

func main() {
	// BEGIN OMIT
	var a, b int = 42, 42
	fmt.Println(a == b)

	var i, j interface{} = a, b
	fmt.Println(i == j)

	var s, t struct{ i interface{} }
	s.i, t.i = a, b
	fmt.Println(s == t)
	// END OMIT
}
