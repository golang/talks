package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("use %s varname\n", os.Args[0])
		os.Exit(1)
	}
	fmt.Println(os.Getenv(os.Args[1]))
}
