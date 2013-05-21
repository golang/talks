package main

import (
	"fmt"
	"time"
)

// STARTMAIN1 OMIT
type Ball struct{}

func main() {
	table := make(chan Ball)
	stop1, stop2 := make(chan int), make(chan int)
	go player("ping", table, stop1)
	go player("pong", table, stop2)

	table <- Ball{}
	time.Sleep(1 * time.Second)
	stop1 <- 1
	stop2 <- 1
	<-stop1
	<-stop2
}

// STOPMAIN1 OMIT

// STARTPLAYER1 OMIT
func player(name string, table chan Ball, stop chan int) {
	for {
		select {
		case ball := <-table:
			fmt.Println(name)
			time.Sleep(100 * time.Millisecond)
			table <- ball
		case <-stop:
			fmt.Println(name, "done")
			stop <- 1
		}
	}
}

// STOPPLAYER1 OMIT
