package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetURLsFromHTML(t *testing.T) {
	htmlBody := `<html>
	<body>
		<a href="/path/one">
			<span>Some page</span>
		</a>
		<a href="https://other.com/path/one">
			<span>some other page</span>
		</a>
	</body>
</html>`
	inputUrl := "https://vikuuu.github.io"

	malformedHtml := `<html>
	<ch<bankai>>
	<body>
</htl`

	// Wrong raw base url provided
	wrongRawURL := "://bankai.com"
	_, err := getURLsFromHTML(htmlBody, wrongRawURL)
	assert.Error(t, err)

	// Malformed html body passed still no error should be returned
	_, err = getURLsFromHTML(malformedHtml, inputUrl)
	assert.NoError(t, err)

	// get the valid links out
	expected := []string{"https://vikuuu.github.io/path/one", "https://other.com/path/one"}
	got, err := getURLsFromHTML(htmlBody, inputUrl)
	assert.NoError(t, err)
	assert.Equal(t, expected, got)

	// malformed url in anchor tag
	htmlBody = `<html>
	<body>
		<a href="://other.com/path/one">
			<span>some other page</span>
		</a>
	</body>
</html>`
	got, err = getURLsFromHTML(htmlBody, inputUrl)
	assert.Zero(t, len(got))
}
