// +build ignore

package strings_test

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	for _, test := range []struct {
		s   string
		sep string
		out int
	}{
		{"", "", 0},
		{"", "a", -1},
		{"fo", "foo", -1},
		{"foo", "foo", 0},
		{"oofofoofooo", "f", 2},
		// etc
	} {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q,%q) = %v; want %v", test.s, test.sep, actual, test.out)
		}
	}
}
