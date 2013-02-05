// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

// playScript registers an HTTP handler at /play.js that contains all the
// scripts specified by path, relative to basePath.
func playScript(path ...string) {
	modTime := time.Now()
	var buf bytes.Buffer
	for _, p := range append(path, "/static/play.js") {
		b, err := ioutil.ReadFile(filepath.Join(basePath, p))
		if err != nil {
			panic(err)
		}
		buf.Write(b)
	}
	b := buf.Bytes()
	http.HandleFunc("/play.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/javascript")
		http.ServeContent(w, r, "", modTime, bytes.NewReader(b))
	})
}
