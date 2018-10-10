// +build ignore

package main

import (
	"fmt"
	"go/build"
)

func main() {
	pkg, _ := build.Import("net/http", "", 0) // HL
	fmt.Println(pkg.Dir)
	fmt.Println(pkg.GoFiles)
}
