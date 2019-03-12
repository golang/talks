// +build ignore,OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/talks/content/2016/applicative/google"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START2 OMIT
func main() {
	start := time.Now()
	search := google.First( // HL
		google.FakeSearch("replica 1", "I'm #1!", ""),  // HL
		google.FakeSearch("replica 2", "#2 wins!", "")) // HL
	result := search("golang")
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

// STOP2 OMIT
