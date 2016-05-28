// +build OMIT

package main

import "fmt"

// START SLICES OMIT
var arr [8]Rectangle

var (
	rects  = arr[2:4]
	rects2 = []Rectangle{rects[0], rects[1]}
)

// START TYPES OMIT
type Point struct {
	X, Y int
}

type Rectangle struct {
	Min, Max Point
}

// END TYPES OMIT

func main() {
	var r0 Rectangle

	r1 := r0 // struct copy

	r1.Min.X, r1.Min.Y = -1, -1
	r1.Max = Point{X: 2}

	fmt.Printf("r0 is %+v\n", r0)
	fmt.Printf("r1 is %v\n", r1)
}
