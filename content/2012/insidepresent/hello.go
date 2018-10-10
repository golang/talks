// +build OMIT

package main

import ( "fmt"; "time" )

func main() {
	for {
		fmt.Println("Hello, Gophers!")
		time.Sleep(time.Second)
	}
}
