package main

import (
	"fmt"
	"time"
)

var list []string

func main() {
	Crawl("http://golang.org/", 4, fetcher)
	// main()->Crawl(url, depth, fakeFetcher)->fakeFetcher(implements Fetcher interface)->Fetch()

	time.Sleep(500 * time.Millisecond)
}

func checkUrlInSlice(url string, list []string) bool  {
	for _, b := range list {
		if b == url {
			return true
		}
	}

	return false
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher fakeFetcher) {
	for k, _ := range fetcher {
		go runner(k, depth, fetcher)
	}
}

func runner(k string, depth int, fetcher fakeFetcher)  {
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}

	isExists := checkUrlInSlice(k, list)
	if isExists {
		return
	} else {
		list = append(list, k)
	}

	body, urls, err := fetcher.Fetch(k)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", k, body)
	for _, u := range urls {
		go runner(u, depth-1, fetcher)
	}
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		fmt.Println("fdsfsfdsfd", res, ok)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
