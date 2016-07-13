package main

import "testing"

func TestLowerThanOneIsNotPrimeNumber(t *testing.T) {
	p := IsPrimeNumber(0)
	if p {
		t.Errorf("Expected %q but got %q", false, p)
	}
}
