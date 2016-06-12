package main

import "testing"

func TestWordWrapWithEmptyString(t *testing.T) {
	w := WordWrap("", 0)
	if w != "" {
		t.Errorf("Expected %q but got %q", "", w)
	}

}

func TestWordWrapWithShorterString(t *testing.T) {
	w := WordWrap("a", 5)
	if w != "a" {
		t.Errorf("Expected %q but got %q", "a", w)
	}
}
