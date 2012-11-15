// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"fmt"
	"html/template"
	"strings"
)

func init() {
	Register("image", parseImage, image)
}

type Image struct {
	File string
	Args []interface{}
}

func (i Image) HTML(t *template.Template) (template.HTML, error) {
	return execTemplate(t, "image", i)
}

func parseImage(fileName string, lineno int, text string) (Elem, error) {
	args := strings.Fields(text)
	a, err := parseArgs(fileName, lineno, args[2:])
	if err != nil {
		return nil, err
	}
	return Image{File: args[1], Args: a}, nil
}

// image is the entry point for the '.image' present command.
func image(file string, arg []interface{}) (template.HTML, error) {
	args := ""
	switch len(arg) {
	case 0:
		// no size parameters
	case 2:
		args = fmt.Sprintf("height='%v' width='%v'", arg[0], arg[1])
	default:
		return "", fmt.Errorf("incorrect image invocation: code %q %v", file, arg)
	}
	return template.HTML(fmt.Sprintf(`<img src=%q %s>`, file, args)), nil
}
