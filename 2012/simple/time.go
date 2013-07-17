// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	if time.Now().Hour() < 12 {
		fmt.Println("Good morning.")
	} else {
		fmt.Println("Good afternoon (or evening).")
	}
	// END OMIT
}
