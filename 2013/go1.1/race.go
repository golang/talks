// +build OMIT

package main

func main() {
	var a int
	go func() {
		for {
			if a == 0 {
				a = 1
			}
		}
	}()
	for {
		if a == 1 {
			a = 0
		}
	}
}
