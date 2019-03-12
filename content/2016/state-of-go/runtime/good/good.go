// +build ignore,OMIT

package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	count(10000) // what if we have 1, 2, 25?
}

func count(n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	m := map[int]int{}
	var mu sync.Mutex // HL
	for i := 1; i <= n; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				mu.Lock() // HL
				m[i]++
				mu.Unlock() // HL
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
}
