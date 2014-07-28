// +build OMIT

package main

import (
	"bytes"
	"net"
)

func handleConn(conn net.Conn) {
	// does something that should be tested.
}

type loopBack struct {
	net.Conn
	buf bytes.Buffer
}

func (c *loopBack) Read(b []byte) (int, error) {
	return c.buf.Read(b)
}

func (c *loopBack) Write(b []byte) (int, error) {
	return c.buf.Write(b)
}
