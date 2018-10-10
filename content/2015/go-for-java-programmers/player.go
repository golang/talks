// +build OMIT

package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{ hits int }

func main() {
	player("ping", new(Ball))
}

func player(name string, ball *Ball) {
	for i := 0; ; i++ {
		ball.hits++
		fmt.Println(name, i, "hit", ball.hits)
		time.Sleep(100 * time.Millisecond)
	}
}

// STOPMAIN1 OMIT
