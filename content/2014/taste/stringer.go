// +build OMIT

package main

import "fmt"

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.x, p.y)
}

type Weekday int

const (
	Mon Weekday = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

var names = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Weekday) String() string { // ...
	return names[d]
}

// Stringer START OMIT
type Stringer interface {
	String() string
}

// Stringer END OMIT

// main START OMIT
func main() {
	var x Stringer
	x = Point{2, 3}
	fmt.Println("A", x.String())

	x = Tue
	fmt.Println("B", x.String())

	fmt.Println("C", Point{2, 3}) // fmt.Println knows about Stringer!
	fmt.Println("D", Tue)
}

// main END OMIT
