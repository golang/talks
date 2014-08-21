// +build OMIT

package main

import "fmt"

// type START OMIT
type Weekday int

// type END OMIT

const (
	Mon Weekday = iota
	Tue
	Wed
	Thu
	Fri
	Sat
	Sun
)

var names = [...]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// String START OMIT
func (d Weekday) String() string { // ...
	// String END OMIT
	return names[d]
}

// main START OMIT
func main() {
	fmt.Println(Mon.String())
	fmt.Println()

	for d := Mon; d <= Sun; d++ {
		fmt.Println(d.String())
	}
}

// main END OMIT
