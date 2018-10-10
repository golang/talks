// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// START1 OMIT
// Search runs query on a backend and returns the result.
type Search func(query string) Result
type Result struct {
	Hit string
	Err error
}

// First runs query on replicas and returns the first result.
func First(query string, replicas ...Search) Result {
	c := make(chan Result, len(replicas))
	search := func(replica Search) { c <- replica(query) }
	for _, replica := range replicas {
		go search(replica)
	}
	return <-c
}

// STOP1 OMIT

// START2 OMIT
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First("golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Printf("%+v\n", result)
	fmt.Println(elapsed)
}

// STOP2 OMIT

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{Hit: fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}
