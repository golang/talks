// +build OMIT

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:4000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	go dial()

	c, err := l.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	io.Copy(os.Stdout, c)
}

func dial() {
	c, err := net.Dial("tcp", "127.0.0.1:4000")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()
	c.Write([]byte("Hello, network\n"))
}
