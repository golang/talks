// +build ignore

package main

import (
	"fmt"
	"runtime"
)

// 1 START OMIT
type P *P
type S []S
type C chan C
type M map[int]M

// 1 END OMIT

// 2 START OMIT
func Val(p *P) int {
	if p == nil {
		return 0
	} else {
		return 1 + Val(*p)
	}
}

func Add(a, b *P) *P {
	if b == nil {
		return a
	} else {
		a1 := new(P)
		*a1 = a // a1 == a + 1
		return Add(a1, *b) // a + b == Add(a+1, b-1)
	}
}

// 2 END OMIT

// 3 START OMIT

func Print(p *P) {
	fmt.Println(Val(p))
}

func Allocate() {
	p := new(P); *p = new(P); **p = new(P)
	runtime.SetFinalizer(p, Print)
	runtime.SetFinalizer(*p, Print)
	runtime.SetFinalizer(**p, Print)
}

func main() {
	Allocate()
	for i := 0; i < 5; i++ {
		runtime.GC()
		runtime.Gosched()
	}
}

// 3 END OMIT
