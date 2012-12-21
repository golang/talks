// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package present

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	var tests = []struct {
		in  string
		out []string
	}{
		{"", []string{}},
		{" ", []string{" "}},
		{"abc", []string{"abc"}},
		{"abc def", []string{"abc", " ", "def"}},
		{"abc def ", []string{"abc", " ", "def", " "}},
		{"hey [[http://golang.org][Gophers]] around",
			[]string{"hey", " ", "[[http://golang.org][Gophers]]", " ", "around"}},
		{"A [[http://golang.org/doc][two words]] link",
			[]string{"A", " ", "[[http://golang.org/doc][two words]]", " ", "link"}},
		{"Visit [[http://golang.org/doc]] now",
			[]string{"Visit", " ", "[[http://golang.org/doc]]", " ", "now"}},
		{"not [[http://golang.org/doc][a [[link]] ]] around",
			[]string{"not", " ", "[[http://golang.org/doc][a [[link]]", " ", "]]", " ", "around"}},
		{"[[http://golang.org][foo bar]]",
			[]string{"[[http://golang.org][foo bar]]"}},
		{"ends with [[http://golang.org][link]]",
			[]string{"ends", " ", "with", " ", "[[http://golang.org][link]]"}},
	}
	for _, test := range tests {
		out := split(test.in)
		if !reflect.DeepEqual(out, test.out) {
			t.Errorf("split(%q):\ngot\t%q\nwant\t%q", test.in, out, test.out)
		}
	}
}

func TestFont(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"", ""},
		{" ", " "},
		{"\tx", "\tx"},
		{"_a_", "<i>a</i>"},
		{"*a*", "<b>a</b>"},
		{"`a`", "<code>a</code>"},
		{"_a_b_", "<i>a b</i>"},
		{"_a__b_", "<i>a_b</i>"},
		{"_a___b_", "<i>a_ b</i>"},
		{"*a**b*?", "<b>a*b</b>?"},
		{"_a_<>_b_.", "<i>a <> b</i>."},
		{"(_a_)", "(<i>a</i>)"},
		{"((_a_), _b_, _c_).", "((<i>a</i>), <i>b</i>, <i>c</i>)."},
		{"(_a)", "(_a)"},
		{"(_a)", "(_a)"},
		{"_Why_use_scoped__ptr_? Use plain ***ptr* instead.", "<i>Why use scoped_ptr</i>? Use plain <b>*ptr</b> instead."},
		{"_hey_ [[http://golang.org][*Gophers*]] *around*",
			`<i>hey</i> <a href="http://golang.org" target="_blank"><b>Gophers</b></a> <b>around</b>`},
		{"_hey_ [[http://golang.org][so _many_ *Gophers*]] *around*",
			`<i>hey</i> <a href="http://golang.org" target="_blank">so <i>many</i> <b>Gophers</b></a> <b>around</b>`},
		{"Visit [[http://golang.org]] now",
			`Visit <a href="http://golang.org" target="_blank">http://golang.org</a> now`},
	}
	for _, test := range tests {
		out := font(test.in)
		if out != test.out {
			t.Errorf("font(%q):\ngot\t%q\nwant\t%q", test.in, out, test.out)
		}
	}
}

func TestStyle(t *testing.T) {
	var tests = []struct {
		in  string
		out string
	}{
		{"", ""},
		{" ", " "},
		{"\tx", "\tx"},
		{"_a_", "<i>a</i>"},
		{"*a*", "<b>a</b>"},
		{"`a`", "<code>a</code>"},
		{"_a_b_", "<i>a b</i>"},
		{"_a__b_", "<i>a_b</i>"},
		{"_a___b_", "<i>a_ b</i>"},
		{"*a**b*?", "<b>a*b</b>?"},
		{"_a_<>_b_.", "<i>a &lt;&gt; b</i>."},
		{"(_a_<>_b_)", "(<i>a &lt;&gt; b</i>)"},
		{"((_a_), _b_, _c_).", "((<i>a</i>), <i>b</i>, <i>c</i>)."},
		{"(_a)", "(_a)"},
	}
	for _, test := range tests {
		out := string(style(test.in))
		if out != test.out {
			t.Errorf("style(%q):\ngot\t%q\nwant\t%q", test.in, out, test.out)
		}
	}
}
