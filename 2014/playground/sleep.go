// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Good night")
	time.Sleep(8 * time.Hour)
	fmt.Println("Good morning")
}
