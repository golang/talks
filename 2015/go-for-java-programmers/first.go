// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

// START1 OMIT
func First(query string, replicas ...Search) Result {
	c := make(chan Result, len(replicas))
	searchReplica := func(i int) { c <- replicas[i](query) }
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

// STOP1 OMIT

func init() {
	rand.Seed(time.Now().UnixNano())
}

// START2 OMIT
func main() {
	start := time.Now()
	result := First("golang", // HL
		fakeSearch("replica 1"), // HL
		fakeSearch("replica 2")) // HL
	elapsed := time.Since(start)
	fmt.Println(result)
	fmt.Println(elapsed)
}

// STOP2 OMIT

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", kind, query))
	}
}
