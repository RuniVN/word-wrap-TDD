package main

import "testing"

func TestWordWrap(t *testing.T) {
	w := WordWrap("abc", 1)

	if w != "a\nb\nc\n" {
		t.Errorf("Expected %q but got %q", "a\nb\nc\n", w)
	}
}
