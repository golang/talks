// +build OMIT

package main

import (
	"fmt"

	"code.google.com/p/go-tour/tree"
)

func Walk(root *tree.Tree) *Walker {
	return &Walker{stack: []*frame{{t: root}}}
}

type Walker struct {
	stack []*frame
}

type frame struct {
	t  *tree.Tree
	pc int
}

func (w *Walker) Next() (int, bool) {
	if len(w.stack) == 0 {
		return 0, false
	}

	// continued next slide ...
	// CUT OMIT
	f := w.stack[len(w.stack)-1]
	if f.pc == 0 {
		f.pc++
		if l := f.t.Left; l != nil {
			w.stack = append(w.stack, &frame{t: l})
			return w.Next()
		}
	}
	if f.pc == 1 {
		f.pc++
		return f.t.Value, true
	}
	if f.pc == 2 {
		f.pc++
		if r := f.t.Right; r != nil {
			w.stack = append(w.stack, &frame{t: r})
			return w.Next()
		}
	}
	w.stack = w.stack[:len(w.stack)-1]
	return w.Next()
}

// STOP OMIT

func Same(t1, t2 *tree.Tree) bool {
	w1, w2 := Walk(t1), Walk(t2)
	for {
		v1, ok1 := w1.Next()
		v2, ok2 := w2.Next()
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
