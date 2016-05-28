// +build OMIT

package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{ hits int }

func main() {
	go player("ping", new(Ball)) // HL
	time.Sleep(1 * time.Second)
}

func player(name string, ball *Ball) {
	for i := 0; ; i++ {
		ball.hits++
		fmt.Println(name, i, "hit", ball.hits)
		time.Sleep(100 * time.Millisecond)
	}
}

// STOPMAIN1 OMIT
