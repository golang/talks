package main

/*

import (
	"fmt"
	"sort"
)

type Person struct {
	Name     string
	AgeYears int
	SSN      int
}

type byName []Person

func (b byName) Len() int           { return len(b) }
func (b byName) Less(i, j int) bool { return b[i].Name < b[j].Name }
func (b byName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type byAge []Person

func (b byAge) Len() int           { return len(b) }
func (b byAge) Less(i, j int) bool { return b[i].AgeYears < b[j].AgeYears }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

type bySSN []Person

func (b bySSN) Len() int           { return len(b) }
func (b bySSN) Less(i, j int) bool { return b[i].SSN < b[j].SSN }
func (b bySSN) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func main() {
	p := []Person{
		{"Alice", 20, 1234},
		{"Bob", 10, 2345},
		{"Carla", 15, 3456},
	}

	sort.Sort(byName(p))
	fmt.Printf("sorted by name: %v\n", p)

	sort.Sort(byAge(p))
	fmt.Printf("sorted by age: %v\n", p)

	sort.Sort(bySSN(p))
	fmt.Printf("sorted by SSN: %v\n", p)
}
*/
