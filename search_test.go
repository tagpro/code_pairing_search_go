package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	expected := "Goodbye World"
	if search() != expected {
		t.Errorf("Expected '%s' but got '%v'", expected, search())
	}
}
