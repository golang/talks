// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	birthday, _ := time.Parse("Jan 2 2006", "Nov 10 2009") // time.Time
	age := time.Since(birthday)                            // time.Duration
	fmt.Printf("Go is %d days old\n", age/(time.Hour*24))
	// END OMIT
}
