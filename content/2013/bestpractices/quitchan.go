// +build ignore,OMIT

package main

import (
	"fmt"
	"net"
	"time"
)

// SEND OMIT
func sendMsg(msg, addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

// BROADCAST OMIT
func broadcastMsg(msg string, addrs []string) error {
	errc := make(chan error)
	quit := make(chan struct{})

	defer close(quit)

	for _, addr := range addrs {
		go func(addr string) {
			select {
			case errc <- sendMsg(msg, addr):
				fmt.Println("done")
			case <-quit:
				fmt.Println("quit")
			}
		}(addr)
	}

	for _ = range addrs {
		if err := <-errc; err != nil {
			return err
		}
	}
	return nil
}

// MAIN OMIT
func main() {
	addr := []string{"localhost:8080", "http://google.com"}
	err := broadcastMsg("hi", addr) // HL

	time.Sleep(time.Second)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("everything went fine")
}
