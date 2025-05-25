package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"golang.org/x/net/html"
)

func getHTML(rawUrl string) (string, error) {
	log.Printf("getting html for %s\n", rawUrl)
	c := http.Client{
		Timeout: 15 * time.Second,
	}
	res, err := c.Get(rawUrl)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	// 400+ Status Code returned
	if res.StatusCode >= 400 {
		return "", errors.New(res.Status)
	}
	// Content-Type is not text/html
	contentType := res.Header.Get("content-type")
	if !strings.Contains(contentType, "text/html") {
		return "", errors.New("non valid content-type")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	log.Println("returning html body ", rawUrl)
	return string(body), nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	log.Printf("getting urls from html for %s\n", rawBaseURL)
	baseUrl, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, err
	}
	htmlReader := strings.NewReader(htmlBody)

	log.Println("parsing the html document")
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return nil, err
	}

	log.Println("traversing the doc to get urls")
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
	log.Println("got the urls")

	return urls, nil
}
