package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, "Goodbye World", search())
}
