// +build ignore,OMIT

package main

import (
	"flag"
	"fmt"
	"sort"
	"sync"
)

func main() {
	n := flag.Int("n", 10, "maximum number to consider")
	flag.Parse()

	type pair struct{ n, c int }
	var pairs []pair
	for n, c := range countFactorsWideSection(*n) {
		pairs = append(pairs, pair{n, c})
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i].n < pairs[j].n })
	for _, p := range pairs {
		fmt.Printf("%3d: %3d\n", p.n, p.c)
	}
}

func countFactorsNarrowSection(n int) map[int]int {
	m := map[int]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(n - 1)
	for i := 2; i <= n; i++ {
		go func(i int) {
			// NARROW OMIT
			for _, f := range factors(i) {
				mu.Lock() // HL
				m[f]++
				mu.Unlock() // HL
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return m
}

func countFactorsWideSection(n int) map[int]int {
	m := map[int]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(n - 1)
	for i := 2; i <= n; i++ {
		go func(i int) {
			// WIDE OMIT
			mu.Lock() // HL
			for _, f := range factors(i) {
				m[f]++
			}
			mu.Unlock() // HL
			wg.Done()
		}(i)
	}
	wg.Wait()
	return m
}

func countFactorsSeq(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i <= n; i++ {
		for _, f := range factors(i) { // HL
			m[f]++ // HL
		} // HL
	}
	return m
}

func factors(v int) []int {
	var fs []int
	for v > 1 {
		for f := 2; f <= v; f++ {
			if v%f == 0 {
				v = v / f
				fs = append(fs, f)
				break
			}
		}
	}
	return fs
}
