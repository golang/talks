// +build OMIT

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"strings"
)

const blob = `Hey there,
fellow gophers!
Have a good day.
`

func old() {
	// STARTold OMIT
	r := bufio.NewReader(strings.NewReader(blob))
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Println(s)
	}
	// STOP OMIT
}

func main() {
	// STARTnew OMIT
	s := bufio.NewScanner(strings.NewReader(blob))
	for s.Scan() {
		fmt.Println(s.Text())
	}
	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
	// STOP OMIT
}
