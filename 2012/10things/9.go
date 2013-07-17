// +build OMIT

package main

import (
	"fmt"
	"time"
	"math/rand"
)

func waiter(i int, block, done chan struct{}) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	fmt.Println(i, "waiting...")
	<-block // HL
	fmt.Println(i, "done!")
	done <- struct{}{}
}

func main() {
	block, done := make(chan struct{}), make(chan struct{})
	for i := 0; i < 4; i++ {
		go waiter(i, block, done)
	}
	time.Sleep(5 * time.Second)
	close(block) // HL
	for i := 0; i < 4; i++ {
		<-done
	}
}
// endmain OMIT
