package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	// mc := flag.Int("mc", 1, "Set the concurrency")
	// mp := flag.Int("mp", 20, "Set the max pages to crawl")
	// flag.Parse()

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	rawUrl := args[0]
	log.Printf("starting crawl of %s\n", rawUrl)
	mc, _ := strconv.Atoi(args[1])
	mp, _ := strconv.Atoi(args[2])

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	cfg := config{
		pages:              make(map[string]int),
		maxPages:           mp,
		baseUrl:            parsedUrl,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, mc),
		wg:                 &sync.WaitGroup{},
	}
	cfg.crawl()
	fmt.Printf("=============================\n"+
		"  REPORT for %s\n"+
		"=============================\n", rawUrl)
	for k, v := range cfg.pages {
		fmt.Printf("Found %d internal links to %s\n", v, k)
	}
}
