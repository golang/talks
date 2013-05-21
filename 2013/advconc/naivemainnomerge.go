package main

import (
	"fmt"
	"math/rand"
	"time"

	. "code.google.com/p/go.talks/2013/reader"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	FakeFetch = true
}

func main() {
	sub := NaiveSubscribe(Fetch("blog.golang.org"))

	// Close the subscription after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", sub.Close())
	})

	// Print the stream.
	for it := range sub.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	panic("show me the stacks")
}
