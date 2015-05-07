// +build OMIT

package main

import "fmt"

func newInt(v int) *int {
	var n = v
	return &n // HL
}

func inc(p *int) {
	*p++ // try removing * // HL
}

func main() {
	p := newInt(3)
	inc(p)
	fmt.Println(p, "points to", *p)
}
