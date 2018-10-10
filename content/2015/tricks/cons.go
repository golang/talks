// +build ignore

package main

import "fmt"

type cons struct {
	car string
	cdr interface{}
}

func (c cons) String() string {
	if c.cdr == nil || c.cdr == (cons{}) {
		return c.car
	}
	return fmt.Sprintf("%v %v", c.car, c.cdr)
}

func main() {
	m := map[cons]string{}
	c := cons{}
	for _, s := range []string{"life?", "with my", "I doing", "What am"} {
		c = cons{s, c}
	}
	m[c] = "No idea."
	fmt.Println(c, m[c])
}
