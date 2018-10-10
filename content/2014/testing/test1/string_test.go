package strings_test

import (
	"strings"
	"testing"
)

func TestIndex(t *testing.T) {
	const s, sep, want = "chicken", "ken", 4
	got := strings.Index(s, sep)
	if got != want {
		t.Errorf("Index(%q,%q) = %v; want %v", s, sep, got, want)
	}
}
