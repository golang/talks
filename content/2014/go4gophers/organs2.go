// +build OMIT

package main

import (
	"fmt"
	"sort"
)

type Organ struct {
	Name   string
	Weight Grams
}

func (o *Organ) String() string { return fmt.Sprintf("%v (%v)", o.Name, o.Weight) }

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

// PART1 OMIT

type Organs []*Organ

func (s Organs) Len() int      { return len(s) }
func (s Organs) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// PART2 OMIT

type ByName struct{ Organs }

func (s ByName) Less(i, j int) bool { return s.Organs[i].Name < s.Organs[j].Name }

type ByWeight struct{ Organs }

func (s ByWeight) Less(i, j int) bool { return s.Organs[i].Weight < s.Organs[j].Weight }

// PART3 OMIT

func main() {
	// START OMIT
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s}) // HL
	printOrgans("Organs by weight", s)

	sort.Sort(ByName{s}) // HL
	printOrgans("Organs by name", s)
	// STOP OMIT
}

func printOrgans(t string, s []*Organ) {
	fmt.Printf("%s:\n", t)
	for _, o := range s {
		fmt.Printf("  %v\n", o)
	}
}
