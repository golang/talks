// +build OMIT

package main

import "fmt"

type Organ struct {
	Name   string
	Weight Grams
}

func (o *Organ) String() string { return fmt.Sprintf("%v (%v)", o.Name, o.Weight) }

type Grams int

func (g Grams) String() string { return fmt.Sprintf("%dg", int(g)) }

func main() {
	s := []*Organ{{"brain", 1340}, {"heart", 290},
		{"liver", 1494}, {"pancreas", 131}, {"spleen", 162}}

	for _, o := range s {
		fmt.Println(o)
	}
}
