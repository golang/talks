package main

import "sort"

func Reverse(data sort.Interface) sort.Interface {
	return &reverse{data}
}

type reverse struct{ sort.Interface }

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i) // HL
}
