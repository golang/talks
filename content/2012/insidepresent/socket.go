// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build OMIT

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"golang.org/x/net/websocket"
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
// It handles transcoding Messages to and from JSON format, and starting
// and killing Processes.
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
	// END OMIT

	// Start and kill Processes and handle errors.
	proc := make(map[string]*Process)
	for {
		select {
		case m := <-in:
			switch m.Kind {
			case "run":
				proc[m.Id].Kill()
				lOut := limiter(in, out)                      // HL
				proc[m.Id] = StartProcess(m.Id, m.Body, lOut) // HL
			case "kill":
				proc[m.Id].Kill()
			}
		case err := <-errc:
			// A encode or decode has failed; bail.
			log.Println(err)
			// Shut down any running processes.
			for _, p := range proc {
				p.Kill()
			}
			return
		}
	}
}

// Process represents a running process.
type Process struct {
	id   string
	out  chan<- *Message
	done chan struct{} // closed when wait completes
	run  *exec.Cmd
}

// StartProcess builds and runs the given program, sending its output
// and end event as Messages on the provided channel.
func StartProcess(id, body string, out chan<- *Message) *Process {
	p := &Process{
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
func (p *Process) Kill() {
	if p == nil {
		return
	}
	p.run.Process.Kill()
	<-p.done
}

// start builds and starts the given program, sends its output to p.out,
// and stores the running *exec.Cmd in the run field.
func (p *Process) start(body string) error {
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
	// END OMIT

	// build x.go, creating x
	dir, file := filepath.Split(src)
	err := p.cmd(dir, "go", "build", "-o", bin, file).Run()
	defer os.Remove(bin)
	if err != nil {
		return err
	}

	// run x
	cmd := p.cmd("", bin)
	if err = cmd.Start(); err != nil {
		return err
	}

	p.run = cmd
	return nil
}

// wait waits for the running process to complete
// and sends its error state to the client.
func (p *Process) wait() {
	defer close(p.done)
	p.end(p.run.Wait())
}

// end sends an "end" message to the client, containing the process id and the
// given error value.
func (p *Process) end(err error) {
	m := &Message{Id: p.id, Kind: "end"}
	if err != nil {
		m.Body = err.Error()
	}
	p.out <- m
}

// cmd builds an *exec.Cmd that writes its standard output and error to the
// Process' output channel.
func (p *Process) cmd(dir string, args ...string) *exec.Cmd {
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

// END OMIT

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
				// Process produced too much output. Kill it.
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
