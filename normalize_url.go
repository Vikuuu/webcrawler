package main

import (
	"log"
	"net/url"
	"strings"
)

func normalizeURL(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	log.Printf("Parsed URL struct: %q\n", parsedURL)

	return parsedURL.Host + strings.TrimRight(parsedURL.Path, "/")
}
