// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 250)
	boom := time.After(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick")
		case <-boom:
			fmt.Println("boom!")
			return
		}
	}
}
