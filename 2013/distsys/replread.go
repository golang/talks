// +build OMIT

package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	F           = 2
	N           = 5
	ReadQuorum  = F + 1
	WriteQuorum = N - F
)

var delay = false

type Server struct {
	mu   sync.Mutex
	data map[string]*Data
}

type Data struct {
	Key   string
	Value string
	Time  time.Time
}

func (srv *Server) Delay() {
	if delay == false {
		return
	}
	time.Sleep(time.Duration(math.Abs(rand.NormFloat64()*1e9 + 0.1e9)))
}

func (srv *Server) Write(req *Data) {
	t0 := time.Now()
	defer func() {
		if delay {
			fmt.Printf("write took %.3f seconds\n", time.Since(t0).Seconds())
		}
	}()

	srv.mu.Lock()
	defer srv.mu.Unlock()
	srv.Delay()

	if srv.data == nil {
		srv.data = make(map[string]*Data)
	}
	if d := srv.data[req.Key]; d == nil || d.Time.Before(req.Time) {
		srv.data[req.Key] = req
	}
}

func (srv *Server) Read(key string) *Data {
	t0 := time.Now()
	defer func() {
		fmt.Printf("read took %.3f seconds\n", time.Since(t0).Seconds())
	}()

	srv.mu.Lock()
	defer srv.mu.Unlock()
	srv.Delay()

	return srv.data[key]
}

func better(x, y *Data) *Data {
	if x == nil {
		return y
	}
	if y == nil || y.Time.Before(x.Time) {
		return x
	}
	return y
}

func Write(req *Data) {
	t0 := time.Now()
	done := make(chan bool, len(servers))

	for _, srv := range servers {
		go func(srv *Server) {
			srv.Write(req)
			done <- true
		}(srv)
	}

	for n := 0; n < WriteQuorum; n++ {
		<-done
	}

	if delay {
		fmt.Printf("write committed at %.3f seconds\n", time.Since(t0).Seconds())
	}
	for n := WriteQuorum; n < N; n++ {
		<-done
	}
	if delay {
		fmt.Printf("all replicas written at %.3f seconds\n", time.Since(t0).Seconds())
	}
}

func Read(key string) {
	t0 := time.Now()
	replies := make(chan *Data, len(servers))

	for _, srv := range servers {
		go func(srv *Server) {
			replies <- srv.Read(key)
		}(srv)
	}

	var d *Data
	for n := 0; n < ReadQuorum; n++ {
		d = better(d, <-replies)
	}

	if delay {
		fmt.Printf("read committed at %.3f seconds\n", time.Since(t0).Seconds())
	}

	for n := ReadQuorum; n < N; n++ {
		<-replies
	}
	if delay {
		fmt.Printf("all replicas read at %.3f seconds\n", time.Since(t0).Seconds())
	}
}

var servers []*Server

func main() {
	servers = make([]*Server, N)
	for i := range servers {
		servers[i] = new(Server)
	}

	rand.Seed(time.Now().UnixNano())

	delay = false
	Write(&Data{"hello", "there", time.Now()})
	time.Sleep(1 * time.Millisecond)

	Write(&Data{"hello", "world", time.Now()})

	delay = true
	Read("hello")
}
