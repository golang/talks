// +build OMIT

package main

import "fmt"

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

	y := &r1.Max.Y // y is a *int // HL
	*y = 5         // HL
	fmt.Println(y, "points to", *y)

	min := &r1.Min // min is a *Point // HL
	min.X = 7      // HL
	fmt.Printf("r1 is %v\n", r1)
}
