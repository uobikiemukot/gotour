package main

import (
	"fmt"
	"sync"
	"time"
)

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

var cachedResult = make(map[string]*fakeResult)
var mux sync.Mutex

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	wg.Add(1)
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if _, ok := cachedResult[url]; ok {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher, wg)
	}
	return
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		mux.Lock()
		cachedResult[url] = &fakeResult{res.body, res.urls}
		mux.Unlock()
		return res.body, res.urls, nil
	}
	mux.Lock()
	cachedResult[url] = &fakeResult{"", nil}
	mux.Unlock()
	return "", nil, fmt.Errorf("not found: %s", url)
}

func main() {
	var wg sync.WaitGroup
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	//time.Sleep(time.Second)
	time.Sleep(time.Millisecond)
	wg.Wait()
}
