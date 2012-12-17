// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

// playScript registers an HTTP handler at /play.js
// that returns some JavaScript that add script tags
// for each of the provided paths to the document.
func playScript(path ...string) {
	var buf bytes.Buffer
	path = append(path, "/static/play.js")
	for _, p := range path {
		p = url.QueryEscape(p)
		// TODO(adg): make this less awful.
		fmt.Fprintf(&buf, `document.write(unescape("%%3Cscript src='%s' type='text/javascript'%%3E%%3C/script%%3E"));`, p)
	}
	b := buf.Bytes()
	http.HandleFunc("/play.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/javascript")
		w.Write(b)
	})
}
