// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Nanosecond)
		fmt.Println(time.Since(start))
	}
}
