// +build OMIT

package main

import (
	"flag"
	"fmt"
	"time"
)

var (
	message = flag.String("message", "Hello!", "what to say")
	delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

func main() {
	flag.Parse()
	fmt.Println(*message)
	time.Sleep(*delay)
}
