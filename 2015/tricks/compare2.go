// +build ignore

package main

import "fmt"

func main() {
	// BEGIN OMIT
	var q, r struct{ s []string }
	fmt.Println(q == r)
	// END OMIT
}
