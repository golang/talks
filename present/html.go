package main

import (
	"errors"
	"html/template"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func init() {
	Register("html", parseHTML, nil)
}

func parseHTML(fileName string, lineno int, text string) (Elem, error) {
	p := strings.Fields(text)
	if len(p) != 2 {
		return nil, errors.New("invalid .html args")
	}
	name := filepath.Join(filepath.Dir(fileName), p[1])
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return HTML(b), nil
}

type HTML string

func (s HTML) HTML(*template.Template) (template.HTML, error) {
	return template.HTML(s), nil
}
