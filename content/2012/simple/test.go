// +build OMIT

package main

import "strings"

import "testing"

func TestToUpper(t *testing.T) {
	in := "loud noises"
	want := "LOUD NOISES"
	got := strings.ToUpper(in)
	if got != want {
		t.Errorf("ToUpper(%v) = %v, want %v", in, got, want)
	}
}

func TestContains(t *testing.T) {
	var tests = []struct {
		str, substr string
		expected    bool
	}{
		{"abc", "bc", true},
		{"abc", "bcd", false},
		{"abc", "", true},
		{"", "a", false},
	}
	for _, ct := range tests {
		if strings.Contains(ct.str, ct.substr) != ct.expected {
			t.Errorf("Contains(%s, %s) = %v, want %v",
				ct.str, ct.substr, !ct.expected, ct.expected)
		}
	}
}
