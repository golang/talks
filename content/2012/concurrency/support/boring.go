// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	boring("boring!")
}

// START OMIT
func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}
// STOP OMIT
