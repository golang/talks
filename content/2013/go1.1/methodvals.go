// +build OMIT

package main

import "os"

func main() {
	var w func([]byte) (int, error)
	w = os.Stdout.Write
	w([]byte("hello!\n"))
}
