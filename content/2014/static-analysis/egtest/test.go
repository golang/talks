package egtest

import (
	"fmt"
	"log"
	"strings"
)

func f() {
	fmt.Printf("%s\n", strings.ToLower("HELLO"))
}

func g() {
	var fmt log.Logger
	fmt.Printf("%s\n", strings.ToLower("HELLO"))
}
