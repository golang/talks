// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"fmt"
	"html/template"
	"net/url"
	"strings"
)

func init() {
	Register("link", parseLink, link)
}

type Link struct {
	URL  *url.URL
	Args []string
}

func (l Link) HTML(t *template.Template) (template.HTML, error) {
	return execTemplate(t, "link", l)
}

func parseLink(fileName string, lineno int, text string) (Elem, error) {
	args := strings.Fields(text)
	url, err := url.Parse(args[1])
	if err != nil {
		return nil, err
	}
	return Link{url, args[2:]}, nil
}

// link is the entry point for the '.link' present command.
func link(url url.URL, arg []string) (template.HTML, error) {
	label := ""
	switch len(arg) {
	case 0:
		scheme := url.Scheme + "://"
		if url.Scheme == "mailto" {
			scheme = "mailto:"
		}
		label = strings.Replace(url.String(), scheme, "", 1)
	default:
		label = strings.Join(arg, " ")
	}
	return template.HTML(fmt.Sprintf(`<a href=%q>%s</a>`, url.String(), label)), nil
}
