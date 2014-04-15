// +build OMIT

package main

import "fmt"

func fib(n int) chan int {
	c := make(chan int) // HL
	go func() {         // HL
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a // HL
		}
		close(c)
	}()
	return c
}

func main() {
	for x := range fib(10) {
		fmt.Println(x)
	}
}
