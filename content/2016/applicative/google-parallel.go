// +build ignore

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

func main() {
	start := time.Now()
	results, err := google.SearchParallel("golang") // HLsearch
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed, err)
}
