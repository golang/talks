// +build OMIT

package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func Walk(t *tree.Tree) {
	if t.Left != nil {
		Walk(t.Left)
	}
	fmt.Println(t.Value)
	if t.Right != nil {
		Walk(t.Right)
	}
}

func main() {
	Walk(tree.New(1))
}
