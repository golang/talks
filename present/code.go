// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var playEnabled = false // switched on when socket.go is included (not on App Engine)

func init() {
	Register("code", parseCode, code)
	Register("play", parseCode, code)
}

type Code struct {
	Play       bool   // runnable code
	File       string // file name to read input from
	Cmd        string // text of input line
	Addr       string // really an address
	Highlight  string // HLxxx marker on end of line.
	Type       string // type extension of file (.go etc.).
	SourceFile string
	SourceLine int
}

func (c Code) HTML(t *template.Template) (template.HTML, error) {
	return execTemplate(t, "code", c)
}

// The input line is a .code or .play entry with a file name and an optional HLfoo marker on the end.
// Anything between the file and HL (if any) is an address expression, which we treat as a string here.
// We pick off the HL first, for easy parsing.
var highlightRE = regexp.MustCompile(`\s+HL([a-zA-Z0-9_]+)?$`)
var codeRE = regexp.MustCompile(`\.(code|play)\s+([^\s]+)(\s+)?(.*)?$`)

func parseCode(fileName string, lineno int, text string) (Elem, error) {
	text = strings.TrimSpace(text)
	// Pull off the HL, if any, from the end of the input line.
	highlight := ""
	if hl := highlightRE.FindStringSubmatchIndex(text); len(hl) == 4 {
		highlight = text[hl[2]:hl[3]]
		text = text[:hl[2]-2]
	}
	// Parse the remaining command line.
	args := codeRE.FindStringSubmatch(text)
	// Arguments:
	// args[0]: whole match
	// args[1]:  .code/.play
	// args[2]: file name
	// args[3]: space, if any, before optional address
	// args[4]: optional address
	if len(args) != 5 {
		return nil, fmt.Errorf("%s:%d: syntax error for .code/.play invocation", fileName, lineno)
	}
	command, file, addr := args[1], args[2], strings.TrimSpace(args[4])

	typ := path.Ext(fileName)
	for len(typ) > 0 && typ[0] == '.' {
		typ = typ[1:]
	}
	return Code{
		Play:       command == "play" && playEnabled,
		File:       file,
		Cmd:        text,
		Addr:       addr,
		Highlight:  highlight,
		Type:       typ,
		SourceFile: fileName,
		SourceLine: lineno}, nil
}

// contents reads a file by name and returns its contents as a byte slice.
func contents(name string) ([]byte, error) {
	file, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// code is the entry point for the '.code' present command.
func code(c Code) (template.HTML, error) {
	filename := filepath.Join(filepath.Dir(c.SourceFile), c.File)
	textBytes, err := contents(filename)
	if err != nil {
		return "", fmt.Errorf("%s:%d: %v", c.SourceFile, c.SourceLine, err)
	}
	lo, hi, err := addrToByteRange(c.Addr, 0, textBytes)
	if err != nil {
		return "", fmt.Errorf("%s:%d: %v", c.SourceFile, c.SourceLine, err)
	}
	// Acme patterns stop mid-line, so run to end of line in both directions.
	for lo > 0 && textBytes[lo-1] != '\n' {
		lo--
	}
	for hi < len(textBytes) {
		hi++
		if textBytes[hi-1] == '\n' {
			break
		}
	}
	text := skipOMIT(textBytes[lo:hi])
	// Replace tabs by spaces, which work better in HTML.
	text = strings.Replace(text, "\t", "    ", -1)
	// Escape the program text for HTML.
	text = template.HTMLEscapeString(text)
	// Highlight and span-wrap lines.
	text = "<pre>" + highlightLines(text, c.Highlight) + "</pre>"
	// Include before and after in a hidden span for playground code.
	if c.Play {
		text = hide(skipOMIT(textBytes[:lo])) + text + hide(skipOMIT(textBytes[hi:]))
	}
	// Include the command as a comment.
	text = fmt.Sprintf("<!--{{%s}}\n-->%s", c.Cmd, text)
	return template.HTML(text), nil
}

// skipOMIT turns text into a string, dropping lines ending with OMIT.
func skipOMIT(text []byte) string {
	lines := strings.SplitAfter(string(text), "\n")
	for k := range lines {
		if strings.HasSuffix(lines[k], "OMIT\n") {
			lines[k] = ""
		}
	}
	return strings.Join(lines, "")
}

func parseArgs(name string, line int, args []string) (res []interface{}, err error) {
	res = make([]interface{}, len(args))
	for i, v := range args {
		if len(v) == 0 {
			return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
		}
		switch v[0] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			n, err := strconv.Atoi(v)
			if err != nil {
				return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
			}
			res[i] = n
		case '/':
			if len(v) < 2 || v[len(v)-1] != '/' {
				return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
			}
			res[i] = v
		case '$':
			res[i] = "$"
		default:
			return nil, fmt.Errorf("%s:%d bad code argument %q", name, line, v)
		}
	}
	return
}

// parseArg returns the integer or string value of the argument and tells which it is.
func parseArg(arg interface{}, file string, max int) (ival int, sval string, isInt bool) {
	switch n := arg.(type) {
	case int:
		if n <= 0 || n > max {
			log.Fatalf("%q:%d is out of range", file, n)
		}
		return n, "", true
	case string:
		return 0, n, false
	}
	log.Fatalf("unrecognized argument %v type %T", arg, arg)
	return
}

// oneLine returns the single line generated by a two-argument code invocation.
func oneLine(file, text string, arg interface{}) (line, before, after string, err error) {
	contentBytes, err := contents(file)
	if err != nil {
		return "", "", "", err
	}
	lines := strings.SplitAfter(string(contentBytes), "\n")
	lineNum, pattern, isInt := parseArg(arg, file, len(lines))
	var n int
	if isInt {
		n = lineNum - 1
	} else {
		n, err = match(file, 0, lines, pattern)
		n -= 1
	}
	if err != nil {
		return "", "", "", err
	}
	return lines[n],
		strings.Join(lines[:n], ""),
		strings.Join(lines[n+1:], ""),
		nil
}

// multipleLines returns the text generated by a three-argument code invocation.
func multipleLines(file string, arg1, arg2 interface{}) (line, before, after string, err error) {
	contentBytes, err := contents(file)
	lines := strings.SplitAfter(string(contentBytes), "\n")
	if err != nil {
		return "", "", "", err
	}
	line1, pattern1, isInt1 := parseArg(arg1, file, len(lines))
	line2, pattern2, isInt2 := parseArg(arg2, file, len(lines))
	if !isInt1 {
		line1, err = match(file, 0, lines, pattern1)
	}
	if !isInt2 {
		line2, err = match(file, line1, lines, pattern2)
	} else if line2 < line1 {
		return "", "", "", fmt.Errorf("lines out of order for %q: %d %d", file, line1, line2)
	}
	if err != nil {
		return "", "", "", err
	}
	for k := line1 - 1; k < line2; k++ {
		if strings.HasSuffix(lines[k], "OMIT\n") {
			lines[k] = ""
		}
	}
	return strings.Join(lines[line1-1:line2], ""),
		strings.Join(lines[:line1-1], ""),
		strings.Join(lines[line2:], ""),
		nil
}

// match identifies the input line that matches the pattern in a code invocation.
// If start>0, match lines starting there rather than at the beginning.
// The return value is 1-indexed.
func match(file string, start int, lines []string, pattern string) (int, error) {
	// $ matches the end of the file.
	if pattern == "$" {
		if len(lines) == 0 {
			return 0, fmt.Errorf("%q: empty file", file)
		}
		return len(lines), nil
	}
	// /regexp/ matches the line that matches the regexp.
	if len(pattern) > 2 && pattern[0] == '/' && pattern[len(pattern)-1] == '/' {
		re, err := regexp.Compile(pattern[1 : len(pattern)-1])
		if err != nil {
			return 0, err
		}
		for i := start; i < len(lines); i++ {
			if re.MatchString(lines[i]) {
				return i + 1, nil
			}
		}
		return 0, fmt.Errorf("%s: no match for %#q", file, pattern)
	}
	return 0, fmt.Errorf("unrecognized pattern: %q", pattern)
}

var hlRE = regexp.MustCompile(`(.+) // HL(.*)$`)

// highlightLines emboldens lines that end with "// HL" and
// wraps any other lines in span tags.
func highlightLines(text, label string) string {
	lines := strings.Split(text, "\n")
	for i, line := range lines {
		m := hlRE.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		line := m[1]
		space := ""
		if j := strings.IndexFunc(line, func(r rune) bool {
			return !unicode.IsSpace(r)
		}); j > 0 {
			space = line[:j]
			line = line[j:]
		}
		if m[2] == "" || m[2] == label {
			lines[i] = space + "<b>" + line + "</b>"
		}
	}
	return strings.Join(lines, "\n")
}

func hide(text string) string {
	return fmt.Sprintf(`<pre style="display: none">%s</pre>`, template.HTMLEscapeString(text))
}

const codifyChar = "`"

var codifyRE = regexp.MustCompile(fmt.Sprintf("%s[^%s]+%s", codifyChar, codifyChar, codifyChar))

func codify(s string) template.HTML {
	s = html.EscapeString(s)
	repl := func(s string) string {
		return "<code>" + s[1:len(s)-1] + "</code>"
	}
	return template.HTML(codifyRE.ReplaceAllStringFunc(s, repl))
}
