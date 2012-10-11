// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	slideTemplate = flag.String("template", "", "alternate slide template file")
	parsers       = make(map[string]func(string, int, string) (Elem, error))

	funcs = template.FuncMap{
		"style": style,
	}
)

// Register binds the named action, which does not being with a period, to the
// specified parser and template function to be invoked when the name, with a
// period, appears in the slide input text.
func Register(name string, parser func(fileName string, lineNumber int, inputLine string) (Elem, error), function interface{}) {
	if len(name) == 0 || name[0] == ';' {
		panic("bad name in Register: " + name)
	}
	funcs[name] = function
	parsers["."+name] = parser
}

// renderSlides reads the slide file, builds its template representation,
// and executes the template, sending output to w.
func renderSlides(w io.Writer, base, slideFile string) error {
	// Read the input and build the slide structure.
	pres, err := parse(slideFile, 0)
	if err != nil {
		return err
	}

	// Locate the template file.
	name := filepath.Join(base, "slide.tmpl")
	if *slideTemplate != "" {
		name = *slideTemplate
	}

	// Read and parse the input.
	tmpl := template.New(name).Funcs(funcs)
	if _, err := tmpl.ParseFiles(name); err != nil {
		return err
	}

	pres.Template = tmpl

	// Execute the template.
	return tmpl.ExecuteTemplate(w, "slides", pres)
}

// Pres represents an entire presentation.
type Pres struct {
	Title      string
	Subtitle   string
	Presenters []Presenter
	Slide      []Slide
	Template   *template.Template
}

// Presenter represents the person who wrote and/or is giving the presentation.
type Presenter struct {
	Name    string `json:"name"`
	Company string `json:"company"`
	Gplus   string `json:"gplus"`
	Twitter string `json:"twitter"`
	WWW     string `json:"www"`
}

// Slide represents a single presentation slide.
type Slide struct {
	Number int
	Title  string
	Elem   []Elem
}

// Elem defines the interface for a slide element.
// That is, something that can render itself in HTML.
type Elem interface {
	HTML(t *template.Template) (template.HTML, error)
}

// execTemplate is a helper to execute a template and return the output as a
// template.HTML value.
func execTemplate(t *template.Template, name string, data interface{}) (template.HTML, error) {
	b := new(bytes.Buffer)
	err := t.ExecuteTemplate(b, name, data)
	if err != nil {
		return "", err
	}
	return template.HTML(b.String()), nil
}

// Text represents an optionally preformatted paragraph.
type Text struct {
	Lines []string
	Pre   bool
}

func (t Text) HTML(tmpl *template.Template) (template.HTML, error) {
	return execTemplate(tmpl, "text", t)
}

// List represents a bulleted list.
type List struct {
	Bullet []string
}

func (l List) HTML(t *template.Template) (template.HTML, error) {
	return execTemplate(t, "list", l)
}

// Lines is a helper for parsing line-based input.
type Lines struct {
	line int // 0 indexed, so has 1-indexed number of last line returned
	text []string
}

func readLines(name string) (*Lines, error) {
	contentBytes, err := contents(name)
	if err != nil {
		return nil, err
	}
	return &Lines{0, strings.Split(string(contentBytes), "\n")}, nil
}

func (l *Lines) next() (text string, ok bool) {
	for {
		if l.line >= len(l.text) {
			return "", false
		}
		text = l.text[l.line]
		l.line++
		// Lines starting with # are comments.
		if len(text) == 0 || text[0] != '#' {
			ok = true
			break
		}
	}
	return
}

func (l *Lines) back() {
	l.line--
}

func (l *Lines) nextNonEmpty() (text string, ok bool) {
	for {
		text, ok = l.next()
		if !ok {
			return
		}
		if len(text) > 0 {
			break
		}
	}
	return
}

// parseMode represents flags for the parse function.
type parseMode int

const (
	// If set, parse only the title and subtitle.
	titlesOnly parseMode = 1
)

// parse parses the presentation in the file specified by name.
func parse(name string, mode parseMode) (*Pres, error) {
	pres := new(Pres)
	lines, err := readLines(name)
	if err != nil {
		return nil, err
	}
	var ok bool
	// First non-empty line starts title.
	pres.Title, ok = lines.nextNonEmpty()
	if !ok {
		return nil, errors.New("no title")
	}
	pres.Subtitle, ok = lines.next()
	if !ok {
		return nil, errors.New("no subtitle")
	}
	if mode&titlesOnly > 0 {
		return pres, nil
	}
	// Presenters
	for {
		var text string
		text, ok = lines.nextNonEmpty()
		if !ok {
			return nil, errors.New("unexpected EOF")
		}
		if strings.HasPrefix(text, "* ") {
			lines.back()
			break
		}
		p := Presenter{Name: text}
		p.Company, ok = lines.next()
		if !ok {
			return nil, errors.New("no company")
		}
		p.Gplus, ok = lines.next()
		if !ok {
			return nil, errors.New("no google+ link")
		}
		p.Twitter, ok = lines.next()
		if !ok {
			return nil, errors.New("no twitter name")
		}
		p.WWW, ok = lines.next()
		if !ok {
			return nil, errors.New("no www link")
		}
		pres.Presenters = append(pres.Presenters, p)
	}
	// Slides
	for i := 0; ; i++ {
		var slide Slide
		slide.Number = i
		// Next non-empty line is title.
		text, ok := lines.nextNonEmpty()
		for ok && text == "" {
			text, ok = lines.next()
		}
		if !ok {
			break
		}
		if !strings.HasPrefix(text, "* ") {
			return nil, fmt.Errorf("%s:%d bad title %q", name, lines.line, text)
		}
		slide.Title = text[2:]
		text, ok = lines.nextNonEmpty()
		for ok && !strings.HasPrefix(text, "* ") {
			var e Elem
			r, _ := utf8.DecodeRuneInString(text)
			switch {
			case unicode.IsSpace(r):
				i := strings.IndexFunc(text, func(r rune) bool {
					return !unicode.IsSpace(r)
				})
				if i < 0 {
					break
				}
				indent := text[:i]
				var s []string
				for ok && (strings.HasPrefix(text, indent) || text == "") {
					if text != "" {
						text = text[i:]
					}
					s = append(s, text)
					text, ok = lines.next()
				}
				lines.back()
				pre := strings.Join(s, "\n")
				pre = strings.TrimRightFunc(pre, unicode.IsSpace)
				e = Text{Lines: []string{pre}, Pre: true}
			case strings.HasPrefix(text, "- "):
				var b []string
				for ok && strings.HasPrefix(text, "- ") {
					b = append(b, text[2:])
					text, ok = lines.next()
				}
				lines.back()
				e = List{Bullet: b}
			case strings.HasPrefix(text, "."):
				args := strings.Fields(text)
				parser := parsers[args[0]]
				if parser == nil {
					return nil, fmt.Errorf("%s:%d: unknown command %q\n", name, lines.line, text)
				}
				e, err = parser(name, lines.line, text)
				if err != nil {
					return nil, err
				}
			default:
				var l []string
				for ok && strings.TrimSpace(text) != "" {
					if text[0] == '.' { // Command breaks text block.
						break
					}
					if strings.HasPrefix(text, `\.`) { // Backslash escapes initial period.
						text = text[1:]
					}
					l = append(l, text)
					text, ok = lines.next()
				}
				if len(l) > 0 {
					e = Text{Lines: l}
				}
			}
			if e != nil {
				slide.Elem = append(slide.Elem, e)
			}
			text, ok = lines.nextNonEmpty()
		}
		if strings.HasPrefix(text, "* ") {
			lines.back()
		}
		pres.Slide = append(pres.Slide, slide)
	}
	return pres, nil
}
