// +build OMIT

package main

// START OMIT
import "fmt"

const digits = "0123456789abcdef"

type Point struct {
	x, y int
	tag  string
}

var s [32]byte

var msgs = []string{"Hello, 世界", "Ciao, Mondo"}

func itoa(x, base int) string

// STOP OMIT

func main() {
	fmt.Println() // use fmt
}
