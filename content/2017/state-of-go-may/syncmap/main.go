// +build go1.9

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var m sync.Map

	for i := 0; i < 3; i++ {
		go func(i int) {
			for j := 0; ; j++ {
				m.Store(i, j)
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		m.Range(func(key, value interface{}) bool {
			fmt.Printf("%d: %d\t", key, value)
			return true
		})
		fmt.Println()
		time.Sleep(time.Second)
	}
}
