// +build OMIT

package main

func f(x int) int {
	return x / 0
}

func main() {
	f(1)
}
