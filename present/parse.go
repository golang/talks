// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	parsers = make(map[string]func(string, int, string) (Elem, error))

	funcs = template.FuncMap{
		"style": style,
	}
)

// Register binds the named action, which does not being with a period, to the
// specified parser and template function to be invoked when the name, with a
// period, appears in the present input text.
// The function argument is an optional template function that is available
// inside templates under that name.
func Register(name string, parser func(fileName string, lineNumber int, inputLine string) (Elem, error), function interface{}) {
	if len(name) == 0 || name[0] == ';' {
		panic("bad name in Register: " + name)
	}
	parsers["."+name] = parser
	if function != nil {
		funcs[name] = function
	}
}

// extensions maps the presentable file extensions to the name of the
// template to be executed.
var extensions = map[string]string{
	".slide":   "slides.tmpl",
	".article": "article.tmpl",
}

func isDoc(path string) bool {
	_, ok := extensions[filepath.Ext(path)]
	return ok
}

// renderDoc reads the present file, builds its template representation,
// and executes the template, sending output to w.
func renderDoc(w io.Writer, base, docFile string) error {
	// Read the input and build the doc structure.
	pres, err := parse(docFile, 0)
	if err != nil {
		return err
	}

	// Find which template should be executed.
	ext := filepath.Ext(docFile)
	contentTmpl, ok := extensions[ext]
	if !ok {
		return fmt.Errorf("no template for extension %v", ext)
	}

	// Locate the template file.
	actionTmpl := filepath.Join(base, "templates/action.tmpl")
	contentTmpl = filepath.Join(base, "templates", contentTmpl)

	// Read and parse the input.
	tmpl := template.New("").Funcs(funcs)
	if _, err := tmpl.ParseFiles(actionTmpl, contentTmpl); err != nil {
		return err
	}

	pres.Template = tmpl

	// Execute the template.
	return tmpl.ExecuteTemplate(w, "root", pres)
}

// Doc represents an entire document.
type Doc struct {
	Title    string
	Subtitle string
	Authors  []Author
	Sections []Section
	Template *template.Template
}

// Author represents the person who wrote and/or is presenting the document.
type Author struct {
	Elem []Elem
}

// TextElem returns the first text elements of the author details.
// This is used to display the author' name, job title, and company
// without the contact details.
func (p *Author) TextElem() (elems []Elem) {
	for _, el := range p.Elem {
		if _, ok := el.(Text); !ok {
			break
		}
		elems = append(elems, el)
	}
	return
}

// Section represents a section of a document (such as a presentation slide)
// comprising a title and a list of elements.
type Section struct {
	Number []int
	Title  string
	Elem   []Elem
	Doc    *Doc
}

func (s Section) Sections() (sections []Section) {
	for _, e := range s.Elem {
		if s, ok := e.(Section); ok {
			sections = append(sections, s)
		}
	}
	return
}

// Level returns the level of the given section.
// The document title is level 1, main section 2, etc.
func (s Section) Level() int {
	return len(s.Number) + 1
}

// FormattedNumber returns a string containing the concatenation of the
// numbers identifying a Section.
func (s Section) FormattedNumber() string {
	b := &bytes.Buffer{}
	for _, n := range s.Number {
		fmt.Fprintf(b, "%v.", n)
	}
	return b.String()
}

func (s Section) HTML(tmpl *template.Template) (template.HTML, error) {
	return execTemplate(tmpl, "section", s)
}

// Elem defines the interface for a present element.
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
	contentBytes, err := ioutil.ReadFile(name)
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

// parse parses the document in the file specified by name.
func parse(name string, mode parseMode) (*Doc, error) {
	doc := new(Doc)
	lines, err := readLines(name)
	if err != nil {
		return nil, err
	}
	var ok bool
	// First non-empty line starts title.
	doc.Title, ok = lines.nextNonEmpty()
	if !ok {
		return nil, errors.New("no title")
	}
	doc.Subtitle, ok = lines.next()
	if !ok {
		return nil, errors.New("no subtitle")
	}
	if mode&titlesOnly > 0 {
		return doc, nil
	}
	// Authors
	if doc.Authors, err = parseAuthors(lines); err != nil {
		return nil, err
	}
	// Sections
	if doc.Sections, err = parseSections(name, lines, []int{}, doc); err != nil {
		return nil, err
	}
	return doc, nil
}

// lesserHeading returns true if text is a heading of a lesser or equal level
// than that denoted by prefix.
func lesserHeading(text, prefix string) bool {
	return strings.HasPrefix(text, "*") && !strings.HasPrefix(text, prefix+"*")
}

// parseSections parses Sections from lines for the section level indicated by
// number (a nil number indicates the top level).
func parseSections(name string, lines *Lines, number []int, doc *Doc) ([]Section, error) {
	var sections []Section
	for i := 1; ; i++ {
		// Next non-empty line is title.
		text, ok := lines.nextNonEmpty()
		for ok && text == "" {
			text, ok = lines.next()
		}
		if !ok {
			break
		}
		prefix := strings.Repeat("*", len(number)+1)
		if !strings.HasPrefix(text, prefix+" ") {
			lines.back()
			break
		}
		section := Section{
			Number: append(append([]int{}, number...), i),
			Title:  text[len(prefix)+1:],
			Doc:    doc,
		}
		text, ok = lines.nextNonEmpty()
		for ok && !lesserHeading(text, prefix) {
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
			case strings.HasPrefix(text, prefix+"* "):
				lines.back()
				subsecs, err := parseSections(name, lines, section.Number, doc)
				if err != nil {
					return nil, err
				}
				for _, ss := range subsecs {
					section.Elem = append(section.Elem, ss)
				}
			case strings.HasPrefix(text, "."):
				args := strings.Fields(text)
				parser := parsers[args[0]]
				if parser == nil {
					return nil, fmt.Errorf("%s:%d: unknown command %q\n", name, lines.line, text)
				}
				t, err := parser(name, lines.line, text)
				if err != nil {
					return nil, err
				}
				e = t
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
				section.Elem = append(section.Elem, e)
			}
			text, ok = lines.nextNonEmpty()
		}
		if strings.HasPrefix(text, "*") {
			lines.back()
		}
		sections = append(sections, section)
	}
	return sections, nil
}

func parseAuthors(lines *Lines) (authors []Author, err error) {
	// This grammar demarcates authors with blanks.

	// Skip blank lines.
	if _, ok := lines.nextNonEmpty(); !ok {
		return nil, errors.New("unexpected EOF")
	}
	lines.back()

	var a *Author
	for {
		text, ok := lines.next()
		if !ok {
			return nil, errors.New("unexpected EOF")
		}

		// If we find a section heading, we're done.
		if strings.HasPrefix(text, "* ") {
			lines.back()
			break
		}

		// If we encounter a blank we're done with this author.
		if a != nil && len(text) == 0 {
			authors = append(authors, *a)
			a = nil
			continue
		}
		if a == nil {
			a = new(Author)
		}

		// Parse the line. Those that
		// - begin with @ are twitter names,
		// - contain slashes are links, or
		// - contain an @ symbol are an email address.
		// The rest is just text.
		var el Elem
		switch {
		case strings.HasPrefix(text, "@"):
			el = parseURL("http://twitter.com/" + text[1:])
			if l, ok := el.(Link); ok {
				l.Args = []string{text}
				el = l
			}
		case strings.Contains(text, ":"):
			el = parseURL(text)
		case strings.Contains(text, "@"):
			el = parseURL("mailto:" + text)
		}
		if el == nil {
			el = Text{Lines: []string{text}}
		}
		a.Elem = append(a.Elem, el)
	}
	if a != nil {
		authors = append(authors, *a)
	}
	return authors, nil
}

func parseURL(text string) Elem {
	u, err := url.Parse(text)
	if err != nil {
		log.Printf("Parse(%q): %v", text, err)
		return nil
	}
	return Link{URL: u}
}
