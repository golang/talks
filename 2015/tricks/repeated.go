// +build ignore

package main

// BEGIN OMIT
type Foo struct {
	i int
	s string
}

var s = []Foo{
	{6 * 9, "Question"},
	{42, "Answer"},
}

var m = map[int]Foo{
	7: {6 * 9, "Question"},
	3: {42, "Answer"},
}

// END OMIT

func main() {}
