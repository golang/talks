// +build OMIT

package main

import (
	"bytes"
	"fmt"
)

type Point struct {
	X, Y int
}

type Rectangle struct {
	Min, Max Point
}

func (r Rectangle) String() string {
	var buf bytes.Buffer
	for h := 0; h < r.Max.Y-r.Min.Y; h++ {
		for w := 0; w < r.Max.X-r.Min.X; w++ {
			buf.WriteString("#")
		}
		buf.WriteString("\n")
	}
	return buf.String()
}

func main() {
	fmt.Println(Rectangle{Max: Point{20, 5}})
}

// EOF OMIT
