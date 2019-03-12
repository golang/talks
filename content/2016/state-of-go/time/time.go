// +build ignore,OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	days := []string{"2015 Feb 29", "2016 Feb 29", "2017 Feb 29"}

	fmt.Println("Are these days valid?")
	for _, day := range days {
		_, err := time.Parse("2006 Jan 2", day)
		fmt.Printf("%v -> %v\n", day, err == nil)
	}
}
