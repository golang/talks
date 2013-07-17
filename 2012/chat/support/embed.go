// +build OMIT

package main

import "fmt"

type A struct{}

func (A) Hello() {
	fmt.Println("Hello!")
}

type B struct {
	A
}

// func (b B) Hello() { b.A.Hello() } // (implicitly!)

func main() {
	var b B
	b.Hello()
}
