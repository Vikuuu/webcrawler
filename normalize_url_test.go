package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeURL(t *testing.T) {
	expected := "blog.boot.dev/path"

	url := "https://blog.boot.dev/path"
	got, err := normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	url = "http://blog.boot.dev/path"
	got, err = normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	url = "http://blog.boot.dev/path/"
	got, err = normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	url = "https://blog.boot.dev/path/"
	got, err = normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	expected = "vikuuu.github.com"
	url = "HTTPS://Vikuuu.github.com"
	got, err = normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	expected = ""
	url = ""
	got, err = normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	expected = "www.github.com/vikuuu"
	url = "http://www.github.com/Vikuuu"
	got, err = normalizeURL(url)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	url = "://bankai.xl"
	got, err = normalizeURL(url)
	assert.Error(t, err)
}
