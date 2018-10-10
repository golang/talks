// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Today is day", time.Now().YearDay())
}
