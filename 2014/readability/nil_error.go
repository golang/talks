// +build OMIT

package main // OMIT

import "log"

type FooError struct{}

func (e *FooError) Error() string { return "foo error" }

func foo() error {
	var ferr *FooError // ferr == nil // HL
	return ferr
}
func main() {
	err := foo()
	if err != nil { // HL
		log.Fatal(err)
	}
}
