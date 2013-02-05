// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package socket implements an WebSocket-based playground backend.
// Clients connect to a websocket handler and send run/kill commands, and
// the server sends the output and exit status of the running processes.
// Multiple clients running multiple processes may be served concurrently.
// The wire format is JSON and is described by the Message type.
//
// This will not run on App Engine as WebSockets are not supported there.
package socket

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"code.google.com/p/go.net/websocket"
)

// Handler implements a WebSocket handler for a client connection.
var Handler = websocket.Handler(socketHandler)

const msgLimit = 1000 // max number of messages to send per session

// Message is the wire format for the websocket connection to the browser.
// It is used for both sending output messages and receiving commands, as
// distinguished by the Kind field.
type Message struct {
	Id   string // client-provided unique id for the process
	Kind string // in: "run", "kill" out: "stdout", "stderr", "end"
	Body string
}

// socketHandler handles the websocket connection for a given present session.
// It handles transcoding Messages to and from JSON format, and starting
// and killing processes.
func socketHandler(c *websocket.Conn) {
	in, out := make(chan *Message), make(chan *Message)
	errc := make(chan error, 1)

	// Decode messages from client and send to the in channel.
	go func() {
		dec := json.NewDecoder(c)
		for {
			var m Message
			if err := dec.Decode(&m); err != nil {
				errc <- err
				return
			}
			in <- &m
		}
	}()

	// Receive messages from the out channel and encode to the client.
	go func() {
		enc := json.NewEncoder(c)
		for m := range out {
			if err := enc.Encode(m); err != nil {
				errc <- err
				return
			}
		}
	}()

	// Start and kill processes and handle errors.
	proc := make(map[string]*process)
	for {
		select {
		case m := <-in:
			switch m.Kind {
			case "run":
				proc[m.Id].Kill()
				lOut := limiter(in, out)
				proc[m.Id] = startProcess(m.Id, m.Body, lOut)
			case "kill":
				proc[m.Id].Kill()
			}
		case err := <-errc:
			if err != io.EOF {
				// A encode or decode has failed; bail.
				log.Println(err)
			}
			// Shut down any running processes.
			for _, p := range proc {
				p.Kill()
			}
			return
		}
	}
}

// process represents a running process.
type process struct {
	id   string
	out  chan<- *Message
	done chan struct{} // closed when wait completes
	run  *exec.Cmd
}

// startProcess builds and runs the given program, sending its output
// and end event as Messages on the provided channel.
func startProcess(id, body string, out chan<- *Message) *process {
	p := &process{
		id:   id,
		out:  out,
		done: make(chan struct{}),
	}
	if err := p.start(body); err != nil {
		p.end(err)
		return nil
	}
	go p.wait()
	return p
}

// Kill stops the process if it is running and waits for it to exit.
func (p *process) Kill() {
	if p == nil {
		return
	}
	p.run.Process.Kill()
	<-p.done // block until process exits
}

// start builds and starts the given program, sending its output to p.out,
// and stores the running *exec.Cmd in the run field.
func (p *process) start(body string) error {
	// We "go build" and then exec the binary so that the
	// resultant *exec.Cmd is a handle to the user's program
	// (rather than the go tool process).
	// This makes Kill work.

	bin := filepath.Join(tmpdir, "compile"+strconv.Itoa(<-uniq))
	src := bin + ".go"
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	// write body to x.go
	defer os.Remove(src)
	err := ioutil.WriteFile(src, []byte(body), 0666)
	if err != nil {
		return err
	}

	// build x.go, creating x
	defer os.Remove(bin)
	dir, file := filepath.Split(src)
	cmd := p.cmd(dir, "go", "build", "-o", bin, file)
	cmd.Stdout = cmd.Stderr // send compiler output to stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	// run x
	cmd = p.cmd("", bin)
	if err := cmd.Start(); err != nil {
		return err
	}
	p.run = cmd
	return nil
}

// wait waits for the running process to complete
// and sends its error state to the client.
func (p *process) wait() {
	p.end(p.run.Wait())
	close(p.done) // unblock waiting Kill calls
}

// end sends an "end" message to the client, containing the process id and the
// given error value.
func (p *process) end(err error) {
	m := &Message{Id: p.id, Kind: "end"}
	if err != nil {
		m.Body = err.Error()
	}
	p.out <- m
}

// cmd builds an *exec.Cmd that writes its standard output and error to the
// process' output channel.
func (p *process) cmd(dir string, args ...string) *exec.Cmd {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Dir = dir
	cmd.Stdout = &messageWriter{p.id, "stdout", p.out}
	cmd.Stderr = &messageWriter{p.id, "stderr", p.out}
	return cmd
}

// messageWriter is an io.Writer that converts all writes to Message sends on
// the out channel with the specified id and kind.
type messageWriter struct {
	id, kind string
	out      chan<- *Message
}

func (w *messageWriter) Write(b []byte) (n int, err error) {
	w.out <- &Message{Id: w.id, Kind: w.kind, Body: string(b)}
	return len(b), nil
}

// limiter returns a channel that wraps dest. Messages sent to the channel are
// sent to dest. After msgLimit Messages have been passed on, a "kill" Message
// is sent to the kill channel, and only "end" messages are passed.
func limiter(kill chan<- *Message, dest chan<- *Message) chan<- *Message {
	ch := make(chan *Message)
	go func() {
		n := 0
		for m := range ch {
			switch {
			case n < msgLimit || m.Kind == "end":
				dest <- m
				if m.Kind == "end" {
					return
				}
			case n == msgLimit:
				// process produced too much output. Kill it.
				kill <- &Message{Id: m.Id, Kind: "kill"}
			}
			n++
		}
	}()
	return ch
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

var uniq = make(chan int) // a source of numbers for naming temporary files

func init() {
	go func() {
		for i := 0; ; i++ {
			uniq <- i
		}
	}()
}
