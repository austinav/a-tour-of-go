# a-tour-of-go
Answers from the go language tour. Decided to create during last exercise so this may not be completed

## Web Crawler
Some potential for creativity over the standard solution with only topics directly covered in the tour
[embedmd]:# (web-crawler.go)
```go
//This was first starting approach, but I tried looking through the go sync doc page
//and found the sync.Map and was skipping right to that without fully understanding how it works
//the real solution uses the pieces discussed in the lesson https://tip.golang.org/tour/solutions/webcrawler.go

//The problem is my solution runs without error, but once we hit the first go Crawl() logging is lost and I am not sure why.
package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var (
	fetched sync.Map
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	fmt.Printf("still print at depth %v\n", depth)
	if depth <= 0 {
		return
	}
	if u, found := fetched.Load(url); found {
		fmt.Printf("%v already fetched\n", u) 
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Printf("%v failed: %v\n", url, err)
		return
	}
	fetched.Store(url, &fakeResult{body, urls})
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		fmt.Printf("%v, %v, %v\n", u, depth-1, fetcher)
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
	
	fmt.Println("\nFetching stats\n--------------")
	fetched.Range(func(key, value any) bool {
		fmt.Printf("%v was fetched\n", key)
		return true
	})
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

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
```
