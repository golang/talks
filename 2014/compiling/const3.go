// +build ignore

package main

import (
	"fmt"
	"unsafe"
)

// 1 START OMIT
var V1 = 0x01020304
var V2 [unsafe.Sizeof(V1)]byte

func main() {
	*(*int)(unsafe.Pointer(&V2)) = V1
	fmt.Println(V2)
}

// 1 END OMIT
