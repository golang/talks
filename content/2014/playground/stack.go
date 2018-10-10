// +build OMIT

package main

func foo(a [1000]byte) {
	foo(a)
}

func main() {
	foo([1000]byte{})
}
