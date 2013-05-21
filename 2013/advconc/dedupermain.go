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

// STARTMAIN OMIT
func main() {
	// STARTMERGECALL OMIT
	// Subscribe to some feeds, and create a merged update stream.
	merged := Dedupe(Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com"))))
	// STOPMERGECALL OMIT

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	panic("show me the stacks")
}

// STOPMAIN OMIT
