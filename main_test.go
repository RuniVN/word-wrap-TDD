package main

import "testing"

func TestWordWrapWithLength1(t *testing.T) {
	w := WordWrap("abc", 1)

	if w != "a\nb\nc\n" {
		t.Errorf("Expected %q but got %q", "a\nb\nc\n", w)
	}

}

func TestWordWrapWithLengthShorterThanWord(t *testing.T) {
	w := WordWrap("hello", 3)

	if w != "hel\nlo" {
		t.Errorf("Expected %q but got %q", "hel\nlo", w)
	}
}
