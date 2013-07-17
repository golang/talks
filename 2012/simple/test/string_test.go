// +build OMIT

package string_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	var tests = []struct {
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
	}
	for _, test := range tests {
		actual := strings.Index(test.s, test.sep)
		if actual != test.out {
			t.Errorf("Index(%q,%q) = %v; want %v", test.s, test.sep, actual, test.out)
		}
	}
}

func BenchmarkIndex(b *testing.B) {
	const s = "some_text=some☺value"
	for i := 0; i < b.N; i++ {
		strings.Index(s, "v")
	}
}

func ExampleIndex() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
	// Output:
	// 4
	// -1
}
