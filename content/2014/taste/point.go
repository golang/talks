// +build OMIT

package main

import "fmt"

// Point START OMIT
type Point struct {
	x, y int
}

// Point END OMIT

// String START OMIT
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

// String END OMIT

// main START OMIT
func main() {
	p := Point{2, 3}
	fmt.Println(p.String())
	fmt.Println(Point{3, 5}.String())
}

// main END OMIT
