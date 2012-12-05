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
	Register("image", parseImage)
}

type Image struct {
	URL        string
	Attributes template.HTML
}

func (i Image) TemplateName() string { return "image" }

func parseImage(fileName string, lineno int, text string) (Elem, error) {
	args := strings.Fields(text)
	img := Image{URL: args[1]}
	a, err := parseArgs(fileName, lineno, args[2:])
	if err != nil {
		return nil, err
	}
	switch len(a) {
	case 0:
		// no size parameters
	case 2:
		attr := fmt.Sprintf(`height="%v" width="%v"`, a[0], a[1])
		img.Attributes = template.HTML(attr)
	default:
		return nil, fmt.Errorf("incorrect image invocation: %q", text)
	}
	return img, nil
}
