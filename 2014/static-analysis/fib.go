// +build OMIT

package main

import "fmt"

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func main() {
	fmt.Println(fib(7))
}
