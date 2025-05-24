package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeURL(t *testing.T) {
	expected := "blog.boot.dev/path"

	url := "https://blog.boot.dev/path"
	got := normalizeURL(url)
	assert.Equal(t, expected, got)

	url = "http://blog.boot.dev/path"
	got = normalizeURL(url)
	assert.Equal(t, expected, got)

	url = "http://blog.boot.dev/path/"
	got = normalizeURL(url)
	assert.Equal(t, expected, got)

	url = "https://blog.boot.dev/path/"
	got = normalizeURL(url)
	assert.Equal(t, expected, got)
}
