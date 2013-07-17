// +build OMIT

package main

// START1 OMIT
var i int
var p, q *Point
var threshold float64 = 0.75

// STOP1 OMIT

// START2 OMIT
var i = 42       // type of i is int
var z = 1 + 2.3i // type of z is complex128
// STOP2 OMIT

func _() int {
	i := 42 // type of i is int
	return &i
}
