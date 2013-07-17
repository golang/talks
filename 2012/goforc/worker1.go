// +build OMIT

package main

import "fmt"
import "time"

// START1 OMIT
func main() {
	start := time.Now()

	in := make(chan int)    // Channel on which work orders are received.
	out := make(chan []int) // Channel on which results are returned.
	go producer(in)
	go worker(in, out) // Launch one worker. // HL
	consumer(out, 100)

	fmt.Println(time.Since(start))
}

// STOP1 OMIT

// START2 OMIT
func worker(in chan int, out chan []int) {
	for {
		order := <-in           // Receive a work order. // HL
		result := factor(order) // Do some work. // HL
		out <- result           // Send the result back. // HL
	}
}

// STOP2 OMIT

// START3 OMIT
func producer(out chan int) {
	for order := 0; ; order++ {
		out <- order // Produce a work order. // HL
	}
}

func consumer(in chan []int, n int) {
	for i := 0; i < n; i++ {
		result := <-in // Consume a result. // HL
		fmt.Println("Consumed", result)
	}
}

// STOP3 OMIT

// factor returns a list containing x and its prime factors.
func factor(x int) (list []int) {
	list = append(list, x)
	for f := 2; x >= f; f++ {
		for x%f == 0 {
			x /= f
			list = append(list, f)
			// Slow down so we can see what happens.
			time.Sleep(50 * time.Millisecond)
		}
	}
	return
}
