// +build OMIT

package main

import "fmt"

// START1 OMIT
const (
	MaxUInt = 1<<64 - 1
	Pi      = 3.14159265358979323846264338327950288419716939937510582097494459
	Pi2     = Pi * Pi
	Delta   = 2.0
)

// STOP1 OMIT

func main() {
	// START2 OMIT
	var x uint64 = MaxUInt
	var pi2 float32 = Pi2
	var delta int = Delta
	// STOP2 OMIT
	fmt.Println("x =", x)
	fmt.Println("pi2 =", pi2)
	fmt.Println("delta =", delta)
}
