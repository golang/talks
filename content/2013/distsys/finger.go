// +build OMIT

package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "serve" {
		serve()
	}
	finger()
}

func finger() {
	c, err := net.Dial("tcp", "localhost:finger")
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(c, "rsc\n")
	io.Copy(os.Stdout, c)
}

func serve() {
	l, err := net.Listen("tcp", "localhost:finger")
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go serveConn(c)
	}
}

func serveConn(c net.Conn) {
	defer c.Close()

	b := bufio.NewReader(c)
	l, err := b.ReadString('\n')
	if err != nil {
		return
	}

	cmd := exec.Command("finger", l[:len(l)-1])
	cmd.Stdout = c
	cmd.Stderr = c
	cmd.Run()
}
