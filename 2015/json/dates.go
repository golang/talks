// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("Standard format: %v\n", now)
	fmt.Printf("American format: %v\n", now.Format("Jan 2 2006"))
	fmt.Printf("European format: %v\n", now.Format("02/01/2006"))
	fmt.Printf("Chinese format: %v\n", now.Format("2006/01/02"))
}
