// +build OMIT

package main // OMIT

import "fmt"

type I interface {
	Key() string
	Value() string
}
type S struct{ I }      // S has method sets of I.
func (s S) Key() string { return "type S" }

func main() {
	var s S
	fmt.Println("key", s.Key())
	fmt.Println(s.Value()) // runtime error: invalid memory address or nil pointer deference  // HL
}
