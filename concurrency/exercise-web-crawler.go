package main

import (
	"fmt"
	"time"
	"sync"
)


type Fetcher interface {
	Fetch(uer string) (body string, urls []string, err error)
}


func Crawl(url string, depth int, fetcher Fetcher) {

	ch := make(chan string, depth * 2)
	var prev int 

	cache := make(map[string]bool)
	var mutex sync.Mutex


	var crawChild func(string, int)
	crawChild = func(url string, depth int) {
		if depth <= 0 {
			return
		}
		if cache[url] == true {
			return
		}
		mutex.Lock()
		defer mutex.Unlock()
		cache[url] = true

		ch <- url
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			go crawChild(u, depth-1)
		}
	}
	crawChild(url, depth)

	for {
        if len(ch) != prev {
            prev = len(ch)
        } else {
            break
        }
        time.Sleep(time.Millisecond)
    }
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

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