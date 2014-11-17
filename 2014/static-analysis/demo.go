// +build OMIT

package main

import "fmt"

type Leaf int

func (l Leaf) Sum() int       { return int(l) }
func (l Leaf) String() string { return fmt.Sprintf("%d", l) }

type Branch struct{ left, rhs Tree }

func (b *Branch) Sum() int       { return b.left.Sum() + b.rhs.Sum() }
func (b *Branch) String() string { return fmt.Sprintf("(%s, %s)", b.left, b.rhs) }

type Tree interface {
	Sum() int
}

func main() {
	var tree Tree = Leaf(42)
	fmt.Println(tree.Sum())

	if unknown {
		tree = &Branch{tree, Leaf(123)}
	}
	fmt.Println(tree.Sum())
	fmt.Println(tree)
}

var unknown bool

//

func _() {
	type Answer struct{ right bool }
	var x struct {
		Answer
		Branch
	}
	fmt.Println(x.right)
}
