// +build OMIT

package main

import (
	"flag"
	"fmt"
)

var message = flag.String("message", "Hello, OSCON!", "what to say")

func main() {
	flag.Parse()
	fmt.Println(*message)
}
