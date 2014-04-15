// +build OMIT

package main

import "fmt"

// prime returns true if n is a prime number.
func prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// fib returns a channel on which the first n Fibonacci numbers are written.
func fib(n int) chan int {
	c := make(chan int)
	go func() {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			a, b = b, a+b
			c <- a
		}
		close(c)
	}()
	return c
}

// filterPrimes returns a channel of ints on which it writes all the prime
// numbers read from cin, and closes the returned channel when cin is closed.
func filterPrimes(cin chan int) chan int {
	cout := make(chan int)
	go func() {
		for v := range cin {
			if prime(v) {
				cout <- v
			}
		}
		close(cout)
	}()
	return cout
}

func main() {
	for p := range filterPrimes(fib(20)) {
		fmt.Println(p)
	}
}
