// +build OMIT

package main

import (
	"fmt"
	"unsafe"
)

type Date struct {
	Day   int
	Month int
	Year  int
}

func main() {
	fmt.Printf("size of %T: %v\n", 0, unsafe.Sizeof(0))
	fmt.Printf("size of %T: %v\n", Date{}, unsafe.Sizeof(Date{}))
	fmt.Printf("size of %T: %v\n", [100]Date{}, unsafe.Sizeof([100]Date{}))
}
