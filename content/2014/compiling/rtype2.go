// +build ignore

package main

import "fmt"

// 1 START OMIT
type F func(*State) F

type State int

func Begin(s *State) F {
	*s = 1
	return Middle
}

func Middle(s *State) F {
	*s++
	if *s >= 10 {
		return End
	}
	return Middle
}
// 1 END OMIT

// 2 START OMIT
func End(s *State) F {
	fmt.Println(*s)
	return nil
}

func main() {
	var f F = Begin
	var s State
	for f != nil {
		f = f(&s)
	}
}

// 2 END OMIT
