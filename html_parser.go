package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}
	htmlReader := strings.NewReader(htmlBody)

	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	urls := []string{}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			attrs := n.Attr
			for _, attr := range attrs {
				if attr.Key == "href" {
					href, err := url.Parse(attr.Val)
					if err != nil {
						fmt.Printf("couldn't parse href '%v': %v\n", attr.Val, err)
						continue
					}
					resolvedUrl := baseUrl.ResolveReference(href)
					urls = append(urls, resolvedUrl.String())
				}
			}
		}
	}

	return urls, nil
}
