package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	result := reverse("Hello, OTUS!")
	expected := "!SUTO ,olleH"

	assert.Equal(t, result, expected, "The two words should be the same.")
}
