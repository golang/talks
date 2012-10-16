// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"

	"code.google.com/p/go.net/websocket"
)

const socketPresent = true

func HandleSocket(path string) {
	http.Handle(path, websocket.Handler(socketHandler))
}

const msgLimit = 1000 // max number of messages to send per session

var uniq = make(chan int) // a source of numbers for naming temporary files

func init() {
	go func() {
		for i := 0; ; i++ {
			uniq <- i
		}
	}()
}

// Message is the wire format for the websocket connection to the browser.
// It is used for both sending output messages and receiving commands, as
// distinguished by the Kind field.
type Message struct {
	Id   string // client-provided unique id for the process
	Kind string // in: "run", "kill" out: "stdout", "stderr", "end"
	Body string
}

// socketHandler handles the websocket connection for a given present session.
// It constructs a new Client and handles transcoding Messages to and from JSON
// format, sending and receiving those messages on the Client's in and out
// channels.
func socketHandler(conn *websocket.Conn) {
	in, out := make(chan *Message), make(chan *Message)
	c := &Client{
		proc: make(map[string]*Process),
		in:   in,
		out:  out,
	}
	go c.loop()

	errc := make(chan error, 1)

	// Decode messages from client and send to the in channel.
	go func() {
		dec := json.NewDecoder(conn)
		for {
			m := new(Message)
			if err := dec.Decode(m); err != nil {
				errc <- err
				close(in)
				return
			}
			in <- m
		}
	}()

	// Receive messages from the out channel and encode to the client.
	go func() {
		enc := json.NewEncoder(conn)
		counts := make(map[string]int)
		for m := range out {
			cnt := counts[m.Id]
			switch {
			case m.Kind == "end" || cnt < msgLimit:
				if err := enc.Encode(m); err != nil {
					errc <- err
					return
				}
				if m.Kind == "end" {
					delete(counts, m.Id)
				}
			case cnt == msgLimit:
				// Process produced too much output. Kill it.
				c.kill(m.Id)
			}
			counts[m.Id]++
		}
	}()

	// Wait for one of the send or receive goroutines to exit.
	if err := <-errc; err != nil && err != io.EOF {
		log.Println(err)
	}

	// Kill any running processes associated with this Client.
	c.Lock()
	for _, p := range c.proc {
		p.kill()
	}
	c.Unlock()
}

// Client represents a connected present client.
// It manages any processes being compiled and run for the client.
type Client struct {
	sync.Mutex // guards proc
	proc       map[string]*Process
	in         <-chan *Message
	out        chan<- *Message
}

// loop handles incoming messages from the client.
func (c *Client) loop() {
	for m := range c.in {
		switch m.Kind {
		case "run":
			c.kill(m.Id)
			go c.run(m.Id, m.Body)
		case "kill":
			c.kill(m.Id)
		}
	}
}

// kill shuts down a running process.
func (c *Client) kill(id string) {
	c.Lock()
	defer c.Unlock()
	if p := c.proc[id]; p != nil {
		p.kill()
	}
}

// run compiles and runs the given program, associating it with the given id.
func (c *Client) run(id, body string) {
	p := NewProcess(id, c.out)
	c.Lock()
	c.proc[id] = p
	c.Unlock()
	err := p.run(body)
	m := &Message{Id: id, Kind: "end"}
	if err != nil {
		m.Body = err.Error()
	}
	c.Lock()
	delete(c.proc, id)
	c.Unlock()
	c.out <- m
}

// Process represents a running process.
type Process struct {
	id             string
	stdout, stderr io.Writer

	sync.Mutex // guards cmd
	cmd        *exec.Cmd
	done       chan struct{} // closed when run complete
}

func NewProcess(id string, out chan<- *Message) *Process {
	return &Process{
		id:     id,
		stdout: newPiper(id, "stdout", out),
		stderr: newPiper(id, "stderr", out),
		done:   make(chan struct{}),
	}
}

// run compiles and runs the given go program.
func (p *Process) run(body string) error {
	defer close(p.done)

	// x is the base name for .go and executable files
	x := filepath.Join(tmpdir, "compile"+strconv.Itoa(<-uniq))
	src := x + ".go"
	bin := x
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	// write body to x.go
	defer os.Remove(src)
	if err := ioutil.WriteFile(src, []byte(body), 0666); err != nil {
		return err
	}

	// build x.go, creating x
	dir, file := filepath.Split(src)
	err := p.exec(dir, "go", "build", "-o", bin, file)
	defer os.Remove(bin)
	if err != nil {
		return err
	}

	// run x
	return p.exec("", bin)
}

// exec runs the specified command in the given directory, writing all standard
// output and standard error to the Process' stdout and stderr fields. It
// stores the running command in the cmd field, and returns when the command
// exits.
func (p *Process) exec(dir string, args ...string) error {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = p.stdout
	cmd.Stderr = p.stderr

	if err := cmd.Start(); err != nil {
		return err
	}

	p.Lock()
	p.cmd = cmd
	p.Unlock()

	err := cmd.Wait()

	p.Lock()
	p.cmd = nil
	p.Unlock()

	return err
}

// kill stops the process if it is running and waits for it to exit.
func (p *Process) kill() {
	p.Lock()
	if p.cmd != nil {
		p.cmd.Process.Kill()
	}
	p.Unlock()
	<-p.done
}

// newPiper returns a writer that converts all writes to Message sends on the
// given channel with the specified id and kind.
func newPiper(id, kind string, out chan<- *Message) io.Writer {
	return &piper{id, kind, out}
}

type piper struct {
	id, kind string
	out      chan<- *Message
}

func (p *piper) Write(b []byte) (n int, err error) {
	p.out <- &Message{
		Id:   p.id,
		Kind: p.kind,
		Body: string(b),
	}
	return len(b), nil
}

var tmpdir string

func init() {
	// find real path to temporary directory
	var err error
	tmpdir, err = filepath.EvalSymlinks(os.TempDir())
	if err != nil {
		log.Fatal(err)
	}
}
