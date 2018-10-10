// +build OMIT

package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func Walk(root *tree.Tree, quit chan struct{}) chan int {
	ch := make(chan int)
	go func() {
		walk(root, ch, quit)
		close(ch)
	}()
	return ch
}

func walk(t *tree.Tree, ch chan int, quit chan struct{}) {
	if t.Left != nil {
		walk(t.Left, ch, quit)
	}
	select { // HL
	case ch <- t.Value: // HL
	case <-quit: // HL
		return // HL
	} // HL
	if t.Right != nil {
		walk(t.Right, ch, quit)
	}
}

// STOP OMIT

func Same(t1, t2 *tree.Tree) bool {
	quit := make(chan struct{}) // HL
	defer close(quit)           // HL
	w1, w2 := Walk(t1, quit), Walk(t2, quit)
	for {
		v1, ok1 := <-w1
		v2, ok2 := <-w2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			return true
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(3), tree.New(3)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
