// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package main

import (
	"code.google.com/p/go.talks/pkg/present"

	_ "code.google.com/p/go.tools/godoc/playground"
)

var basePath = "./present/"

func init() {
	playScript(basePath, "HTTPTransport")
	present.PlayEnabled = true
}
