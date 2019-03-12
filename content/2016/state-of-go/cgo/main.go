// +build ignore,OMIT

package main

/*
int fn(void* arg) { return arg == 0; }
*/
import "C"
import "unsafe"

type T struct{ a, b int }
type X struct{ t *T }

func main() {
	t := T{a: 1, b: 2}
	C.fn(unsafe.Pointer(&t)) // correct // HL

	x := X{t: &t}
	C.fn(unsafe.Pointer(&x)) // incorrect // HL
}
