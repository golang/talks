// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func worker(i int, ch chan Work, quit chan struct{}) {
	for {
		select {
		case w := <-ch:
			if quit == nil { // HL
				w.Refuse(); fmt.Println("worker", i, "refused", w)
				break
			}
			w.Do(); fmt.Println("worker", i, "processed", w)
		case <-quit:
			fmt.Println("worker", i, "quitting")
			quit = nil // HL
		}
	}
}

func main() {
	ch, quit := make(chan Work), make(chan struct{})
	go makeWork(ch)
	for i := 0; i < 4; i++ { go worker(i, ch, quit) }
	time.Sleep(5 * time.Second)
	close(quit)
	time.Sleep(2 * time.Second)
}
// endmain OMIT

type Work string
func (w Work) Do() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}
func (w Work) Refuse() {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
}

func makeWork(ch chan Work) {
	for i := 0; ; i++ {
		ch <- Work(fmt.Sprintf("job %x", i))
	}
}
