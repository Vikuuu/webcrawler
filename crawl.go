package main

import (
	"log"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseUrl            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	_, ok := cfg.pages[normalizedURL]
	if !ok {
		cfg.pages[normalizedURL]++
		isFirst = true
		return isFirst
	}

	return isFirst
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	log.Printf("crawling current url %s\n", rawCurrentURL)
	defer cfg.wg.Done()
	defer func() { <-cfg.concurrencyControl }()

	if len(cfg.pages) >= cfg.maxPages {
		log.Println("max pages reached; stopping the crawler")
		return
	}

	parsedCurrUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	if cfg.baseUrl.Host != parsedCurrUrl.Host {
		return
	}
	nrmlRawCurrentUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("err normalizing url %s %s\n", rawCurrentURL, err)
	}

	isFirst := cfg.addPageVisit(nrmlRawCurrentUrl)
	// we have crawled the page, return
	if !isFirst {
		return
	}

	body, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("%s\n", err)
	}
	urls, err := getURLsFromHTML(body, rawCurrentURL)
	if err != nil {
		log.Printf("%s\n", err)
	}

	for _, url := range urls {
		cfg.wg.Add(1)
		go func(c string) {
			cfg.concurrencyControl <- struct{}{}
			cfg.crawlPage(c)
		}(url)
	}
}

func (cfg *config) crawl() {
	cfg.wg.Add(1)
	go func(url string) {
		cfg.concurrencyControl <- struct{}{}
		cfg.crawlPage(url)
	}(cfg.baseUrl.String())

	cfg.wg.Wait()
	return
}
