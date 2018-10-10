// +build OMIT

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func CheckPassword(c net.Conn) error {
	// read a password from the connection
	buf := make([]byte, 256)
	n, err := c.Read(buf)
	if err != nil {
		return fmt.Errorf("read: %v", err)
	}

	// check it's correct
	got := string(buf[:n])
	if got != "password" {
		return fmt.Errorf("wrong password")
	}
	return nil
}

type fakeConn struct {
	net.Conn
	r io.Reader
}

func (c fakeConn) Read(b []byte) (int, error) {
	return c.r.Read(b)
}

// end_fake OMIT

func main() {
	c := fakeConn{
		r: strings.NewReader("foo"),
	}
	err := CheckPassword(c)
	if err == nil {
		log.Println("expected error using wrong password")
	} else {
		log.Println("OK")
	}
}
