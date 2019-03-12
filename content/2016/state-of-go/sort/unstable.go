// +build ignore,OMIT

package main

import (
	"fmt"
	"sort"
	"strings"
)

type byLength []string

func (b byLength) Len() int           { return len(b) }
func (b byLength) Less(i, j int) bool { return len(b[i]) < len(b[j]) }
func (b byLength) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func main() {
	values := []string{"ball", "hell", "one", "joke", "fool", "moon", "two"}
	sort.Sort(byLength(values))
	fmt.Println(strings.Join(values, "\n"))
}
