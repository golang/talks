// Copyright 2013 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build !appengine

package main

import (
	"net/http"

	"code.google.com/p/go.talks/pkg/socket"
)

// HandleSocket registers the websocket handler with http.DefaultServeMux under
// the given path.
func HandleSocket(path string) {
	playScript("/static/socket.js")
	http.Handle(path, socket.Handler)
}
