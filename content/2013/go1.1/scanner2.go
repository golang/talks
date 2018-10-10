// +build OMIT

package main

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

func main() {
	// START OMIT
	const input = "Now is the winter of our discontent..."
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords) // HL
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	// STOP OMIT
}
