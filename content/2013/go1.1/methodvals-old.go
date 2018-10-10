// +build OMIT

package main

import "os"

func main() {
	var w func([]byte) (int, error)
	w = func(b []byte) (int, error) { return os.Stdout.Write(b) }
	w([]byte("hello!\n"))
}
