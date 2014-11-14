package main

import (
	"fmt"
	"os"
)

const moved = `

The present tool has moved to the Go tools repository.

Please install it from its new location:

	go get golang.org/x/tools/cmd/present


`

func main() {
	fmt.Print(moved)
	os.Exit(1)
}
