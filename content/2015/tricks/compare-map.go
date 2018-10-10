// +build ignore

package main

import "io/ioutil"

func main() {
	// BEGIN OMIT
	a := map[int]bool{}
	a[42] = true

	type T struct {
		i int
		s string
	}

	b := map[*T]bool{}
	b[&T{}] = true

	c := map[T]bool{}
	c[T{37, "hello!"}] = true

	d := map[interface{}]bool{}
	d[42] = true
	d[&T{}] = true
	d[T{123, "four five six"}] = true
	d[ioutil.Discard] = true
	// END OMIT
}
