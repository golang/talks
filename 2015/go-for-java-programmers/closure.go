// +build OMIT

package main

import "fmt"

func div(n, d int) (q, r int, err error) {
	if d == 0 {
		err = fmt.Errorf("%d/%d: divide by zero", n, d)
		return
	}
	return n / d, n % d, nil
}

func main() {
	var failures int

	f := func(n, d int) { // HL
		if q, r, err := div(n, d); err != nil {
			fmt.Println(err)
			failures++ // HL
		} else {
			fmt.Printf("%d/%d = %d leaving %d\n", n, d, q, r)
		}
	}

	f(4, 3)
	f(3, 0)
	fmt.Println("failures:", failures)
}
