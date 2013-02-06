// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"appengine"
	"appengine/urlfetch"

	"code.google.com/p/go.talks/pkg/present"
)

const runURL = "http://golang.org/compile"

var basePath = "./present/"

func init() {
	playScript(basePath, "jquery.js", "playground.js")
	present.PlayEnabled = true
	http.HandleFunc("/compile", compile)
}

func compile(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	b := new(bytes.Buffer)
	if err := passThru(c, b, r); err != nil {
		http.Error(w, "Compile server error.", http.StatusInternalServerError)
		c.Errorf("passThru: %v", err)
		return
	}
	io.Copy(w, b)
}

func passThru(c appengine.Context, w io.Writer, req *http.Request) error {
	client := urlfetch.Client(c)
	defer req.Body.Close()
	r, err := client.Post(runURL, req.Header.Get("Content-type"), req.Body)
	if err != nil {
		return fmt.Errorf("making POST request: %v", err)
	}
	defer r.Body.Close()
	if _, err := io.Copy(w, r.Body); err != nil {
		return fmt.Errorf("copying response Body: %v", err)
	}
	return nil
}
