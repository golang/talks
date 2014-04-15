// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func lookup() {
	const max = 2

	done := make(chan bool, len(worklist))
	limit := make(chan bool, max)

	for _, w := range worklist {
		go func(w *Work) {
			limit <- true
			w.addrs, w.err = LookupHost(w.host)
			<-limit
			done <- true
		}(w)
	}

	for i := 0; i < len(worklist); i++ {
		<-done
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	t0 := time.Now()
	lookup()

	fmt.Printf("\n")
	for _, w := range worklist {
		if w.err != nil {
			fmt.Printf("%s: error: %v\n", w.host, w.err)
			continue
		}
		fmt.Printf("%s: %v\n", w.host, w.addrs)
	}
	fmt.Printf("total lookup time: %.3f seconds\n", time.Since(t0).Seconds())
}

var worklist = []*Work{
	{host: "fast.com"},
	{host: "slow.com"},
	{host: "fast.missing.com"},
	{host: "slow.missing.com"},
}

type Work struct {
	host  string
	addrs []string
	err   error
}

func LookupHost(name string) (addrs []string, err error) {
	t0 := time.Now()
	defer func() {
		fmt.Printf("lookup %s: %.3f seconds\n", name, time.Since(t0).Seconds())
	}()
	h := hosts[name]
	if h == nil {
		h = failure
	}
	return h(name)
}

type resolver func(string) ([]string, error)

var hosts = map[string]resolver{
	"fast.com":         delay(10*time.Millisecond, fixedAddrs("10.0.0.1")),
	"slow.com":         delay(2*time.Second, fixedAddrs("10.0.0.4")),
	"fast.missing.com": delay(10*time.Millisecond, failure),
	"slow.missing.com": delay(2*time.Second, failure),
}

func fixedAddrs(addrs ...string) resolver {
	return func(string) ([]string, error) {
		return addrs, nil
	}
}

func delay(d time.Duration, f resolver) resolver {
	return func(name string) ([]string, error) {
		time.Sleep(d/2 + time.Duration(rand.Int63n(int64(d/2))))
		return f(name)
	}
}

func failure(name string) ([]string, error) {
	return nil, fmt.Errorf("unknown host %v", name)
}
