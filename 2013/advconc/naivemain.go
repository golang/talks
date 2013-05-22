package main

import (
	"fmt"
	"math/rand"
	"time"

	. "code.google.com/p/go.talks/2013/advconc/reader"
)

func NaiveSubscribe(fetcher Fetcher) Subscription {
	s := &naiveSub{
		fetcher: fetcher,
		updates: make(chan Item),
	}
	go s.loop()
	return s
}

type naiveSub struct {
	fetcher Fetcher
	updates chan Item
	closed  bool
	err     error
}

func (s *naiveSub) Updates() <-chan Item {
	return s.updates
}

func (s *naiveSub) loop() {
	// STARTNAIVE OMIT
	for {
		if s.closed { // HLsync
			close(s.updates)
			return
		}
		items, next, err := s.fetcher.Fetch()
		if err != nil {
			s.err = err                  // HLsync
			time.Sleep(10 * time.Second) // HLsleep
			continue
		}
		for _, item := range items {
			s.updates <- item // HLsend
		}
		if now := time.Now(); next.After(now) {
			time.Sleep(next.Sub(now)) // HLsleep
		}
	}
	// STOPNAIVE OMIT
}

func (s *naiveSub) Close() error {
	s.closed = true // HLsync
	return s.err    // HLsync
}

func init() {
	rand.Seed(time.Now().UnixNano())
	FakeFetch = true
}

func main() {
	// Subscribe to some feeds, and create a merged update stream.
	merged := Merge(
		NaiveSubscribe(Fetch("blog.golang.org")),
		NaiveSubscribe(Fetch("googleblog.blogspot.com")),
		NaiveSubscribe(Fetch("googledevelopers.blogspot.com")))

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	// The loops are still running.  Let the race detector notice.
	time.Sleep(1 * time.Second)

	panic("show me the stacks")
}
