package main

import "testing"

func TestWordWrapWithEmptyString(t *testing.T) {
	w := WordWrap("", 0)
	if w != "" {
		t.Errorf("Expected %q but got %q", "", w)
	}

}
