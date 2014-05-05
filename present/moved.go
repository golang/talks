package main

import (
	"fmt"
	"os"
)

const moved = `

The present tool has moved to the go.tools repository.

Please install it from its new location:

	go get code.google.com/p/go.tools/cmd/present


`

func main() {
	fmt.Print(moved)
	os.Exit(1)
}
