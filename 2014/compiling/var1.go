// +build ignore

package main

import (
	"fmt"
	"os"
)

// 1 START OMIT
var V = struct {
	name string
	os.FileMode
}{
	name: "hello.go",
}

func main() {
	fmt.Println(V)
}
// 1 END OMIT
