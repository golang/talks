// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

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
		{"_Why_use_scoped__ptr_? Use plain ***ptr* instead.", "<i>Why use scoped_ptr</i>? Use plain <b>*ptr</b> instead."},
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
	}
	for _, test := range tests {
		out := string(style(test.in))
		if out != test.out {
			t.Errorf("style(%q):\ngot\t%q\nwant\t%q", test.in, out, test.out)
		}
	}
}
