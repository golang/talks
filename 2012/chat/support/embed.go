package main

import "fmt"

type A struct {
	B
}

type B struct{}

func (b B) String() string {
	return "B comes after A"
}

func main() {
	var a A
	fmt.Println(a) // Println calls String to format a
}
