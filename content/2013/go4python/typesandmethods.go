// +build OMIT

package main

import (
	"fmt"
	"strings"
)

type Name struct {
	First  string
	Middle string
	Last   string
}

func (n Name) String() string {
	return fmt.Sprintf("%s %c. %s", n.First, n.Middle[0], strings.ToUpper(n.Last))
}

type SimpleName string

func (s SimpleName) String() string { return string(s) }

func main() {
	n := Name{"William", "Mike", "Smith"}
	fmt.Printf("%s", n.String())
	return
	// second OMIT
	n = Name{"William", "Mike", "Smith"}
	fmt.Println(n)
}
