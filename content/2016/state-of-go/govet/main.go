// +build ignore,OMIT

package main

import "fmt"

func foo() string { return "bar" }

func main() {
	fmt.Printf("%v", foo)
}
