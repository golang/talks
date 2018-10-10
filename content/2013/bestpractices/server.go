// +build ignore,OMIT

package main

import (
	"fmt"
	"time"
)

// START OMIT
type Server struct{ quit chan bool }

func NewServer() *Server {
	s := &Server{make(chan bool)}
	go s.run()
	return s
}

func (s *Server) run() {
	for {
		select {
		case <-s.quit:
			fmt.Println("finishing task")
			time.Sleep(time.Second)
			fmt.Println("task done")
			s.quit <- true
			return
		case <-time.After(time.Second):
			fmt.Println("running task")
		}
	}
}

// STOP OMIT
func (s *Server) Stop() {
	fmt.Println("server stopping")
	s.quit <- true
	<-s.quit
	fmt.Println("server stopped")
}

func main() {
	s := NewServer()
	time.Sleep(2 * time.Second)
	s.Stop()
}
