// +build OMIT

package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func Walk(root *tree.Tree) chan int {
	ch := make(chan int)
	go func() {
		walk(root, ch)
		close(ch)
	}()
	return ch
}

func walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		walk(t.Right, ch)
	}
}

// STOP OMIT

func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := Walk(t1), Walk(t2)
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
