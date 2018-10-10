// +build ignore

package main

// BEGIN1 OMIT
var s = []struct {
	i int
	s string
}{
	struct {
		i int
		s string
	}{6 * 9, "Question"},
	struct {
		i int
		s string
	}{42, "Answer"},
}

var t = []struct {
	i int
	s string
}{
	{6 * 9, "Question"},
	{42, "Answer"},
}

// END OMIT

func main() {}
