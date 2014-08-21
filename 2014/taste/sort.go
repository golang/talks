// +build OMIT

package main

import (
	"fmt"
	"sort"
)

type Weekday int

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

func (d Weekday) String() string { // ...
	return names[d]
}

// lexical START OMIT
type lexical []string

func (a lexical) Len() int           { return len(a) }
func (a lexical) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a lexical) Less(i, j int) bool { return a[i] < a[j] }

// lexical END OMIT

func main() {
	var list []string
	for d := Mon; d <= Sun; d++ {
		list = append(list, d.String())
	}

	sort.Sort(lexical(list))

	for i, x := range list {
		fmt.Println(i, x)
	}
}
