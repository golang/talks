// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
)

// dirList scans the given path and writes a directory listing to w.
// It parses the first part of each .slide file it encounters to display the
// presentation title in the listing.
// If the given path is not a directory, it returns (isDir == false, err == nil)
// and writes nothing to w.
func dirList(w io.Writer, name string) (isDir bool, err error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false, err
	}
	if isDir = fi.IsDir(); !isDir {
		return false, nil
	}
	fis, err := f.Readdir(0)
	if err != nil {
		return false, err
	}
	d := &dirListData{Path: name}
	for _, fi := range fis {
		e := dirListEntry{
			Name: fi.Name(),
			Path: filepath.Join(name, fi.Name()),
		}
		if fi.IsDir() {
			d.Dirs = append(d.Dirs, e)
			continue
		}
		if filepath.Ext(e.Name) == ".slide" {
			if p, err := parse(e.Path, titlesOnly); err != nil {
				log.Println(err)
			} else {
				e.Title = p.Title
			}
			d.Slides = append(d.Slides, e)
		} else {
			d.Other = append(d.Other, e)
		}
	}
	if d.Path == "." {
		d.Path = ""
	}
	return true, dirListTemplate.Execute(w, d)
}

type dirListData struct {
	Path                string
	Dirs, Slides, Other []dirListEntry
}

type dirListEntry struct {
	Name, Path, Title string
}

var dirListTemplate = template.Must(template.New("").Parse(dirListHTML))

const dirListHTML = `<!DOCTYPE html>
<html>
  <head>
    <title>{{.Path}}</title>
    <meta charset='utf-8'>
  </head>
  <body>

  {{with .Path}}<h1>{{.}}</h1>{{end}}

  {{with .Slides}}
  <h2>Slide decks:</h2>
  <ul>
  {{range .}}
  <li><a href="/{{.Path}}">{{.Name}}</a>: {{.Title}}</li>
  {{end}}
  </ul>
  {{end}}

  {{with .Dirs}}
  <h2>Sub-directories:</h2>
  <ul>
  {{range .}}
  <li><a href="/{{.Path}}">{{.Name}}</a></li>
  {{end}}
  </ul>
  {{end}}

  {{with .Other}}
  <h2>Other files:</h2>
  <ul>
  {{range .}}
  <li><a href="/{{.Path}}">{{.Name}}</a></li>
  {{end}}
  </ul>
  {{end}}

  </body>
</html>`
