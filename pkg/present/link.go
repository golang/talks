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
	return template.HTML(renderLink(url.String(), label)), nil
}

func renderLink(url, text string) string {
	text = font(text)
	if text == "" {
		text = url
	}
	return fmt.Sprintf(`<a href="%s" target="_blank">%s</a>`, url, text)
}

// parseInlineLink parses an inline link at the start of s, and returns
// a rendered HTML link and the total length of the raw inline link.
// If no inline link is present, it returns all zeroes.
func parseInlineLink(s string) (link string, length int) {
	if len(s) < 2 || s[:2] != "[[" {
		return
	}
	end := strings.Index(s, "]]")
	if end == -1 {
		return
	}
	urlEnd := strings.Index(s, "]")
	url := s[2:urlEnd]
	const badURLChars = `<>"{}|\^~[] ` + "`" // per RFC1738 section 2.2
	if strings.ContainsAny(url, badURLChars) {
		return
	}
	if urlEnd == end {
		return renderLink(url, ""), end + 2
	}
	if s[urlEnd:urlEnd+2] != "][" {
		return
	}
	text := s[urlEnd+2 : end]
	return renderLink(url, text), end + 2
}
