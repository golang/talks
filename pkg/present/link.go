// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"fmt"
	"net/url"
	"strings"
)

func init() {
	Register("link", parseLink)
}

type Link struct {
	URL   *url.URL
	Label string
}

func (l Link) TemplateName() string { return "link" }

func parseLink(fileName string, lineno int, text string) (Elem, error) {
	args := strings.Fields(text)
	url, err := url.Parse(args[1])
	if err != nil {
		return nil, err
	}
	label := ""
	if len(args) > 2 {
		label = strings.Join(args[2:], " ")
	} else {
		scheme := url.Scheme + "://"
		if url.Scheme == "mailto" {
			scheme = "mailto:"
		}
		label = strings.Replace(url.String(), scheme, "", 1)
	}
	return Link{url, label}, nil
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
