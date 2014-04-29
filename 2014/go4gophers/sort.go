// +build OMIT

package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	// START OMIT
	s := []int{7, 5, 3, 11, 2}
	sort.Sort(IntSlice(s))
	fmt.Println(s)
	// STOP OMIT
}
