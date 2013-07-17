// +build OMIT

package main

import "fmt"

const digits = "0123456789abcdef"

func itoa(x, base int) string {
	// START OMIT
	t := x
	switch {
	case x == 0:
		return "0"
	case x < 0:
		t = -x
	}
	var s [32]byte
	i := len(s)
	for t != 0 { // Look, ma, no ()'s!
		i--
		s[i] = digits[t%base]
		t /= base
	}
	if x < 0 {
		i--
		s[i] = '-'
	}
	return string(s[i:])
	// STOP OMIT
}

func main() {
	fmt.Println(itoa(-42, 2))
}
