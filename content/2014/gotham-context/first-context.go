// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"
)

type Result struct {
	Hit string
	Err error
}

// START1 OMIT
// Search runs query on a backend and returns the result.
type Search func(ctx context.Context, query string) Result // HL

// First runs query on replicas and returns the first result.
func First(ctx context.Context, query string, replicas ...Search) Result {
	c := make(chan Result, len(replicas))
	ctx, cancel := context.WithCancel(ctx)                      // HL
	defer cancel()                                              // HL
	search := func(replica Search) { c <- replica(ctx, query) } // HL
	for _, replica := range replicas {
		go search(replica)
	}
	select {
	case <-ctx.Done(): // HL
		return Result{Err: ctx.Err()} // HL
	case r := <-c:
		return r
	}
}

// STOP1 OMIT

// START2 OMIT
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	result := First(context.Background(),
		"golang",
		fakeSearch("replica 1"),
		fakeSearch("replica 2"))
	elapsed := time.Since(start)
	fmt.Printf("%+v\n", result)
	fmt.Println(elapsed)
}

// STOP2 OMIT

func fakeSearch(kind string) Search {
	return func(ctx context.Context, query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result{Hit: fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}
