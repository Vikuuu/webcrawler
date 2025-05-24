package main

import (
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	links := []string{}
	u, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}
	htmlReader := strings.NewReader(htmlBody)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			attrs := n.Attr
			for _, attr := range attrs {
				if attr.Key == "href" {
					href := attr.Val
					rel, err := u.Parse(href)
					if err != nil {
						return links, err
					}
					links = append(links, rel.String())
				}
			}
		}
	}

	return links, nil
}
