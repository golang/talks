// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package main

const socketPresent = false // no websockets or compilation on app engine (yet)

var basePath = "./present/"

func HandleSocket(path string) {
	panic("websockets not supported on app engine")
}
