package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	expected := "Intended to be broken, please fix the tests..."
	if search() != expected {
		t.Errorf("Expected '%s' but got '%v'", expected, search())
	}
}
