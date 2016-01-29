/// +build OMIT

package main

import (
	"fmt"
	"sync"
)

func main() {
	const workers = 100 // what if we have 1, 2, 25?

	var wg sync.WaitGroup
	wg.Add(workers)
	m := map[int]int{}
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := 0; j < i; j++ {
				m[i]++ // HL
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(m)
}
