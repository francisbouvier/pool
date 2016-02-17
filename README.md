# Pool

A small library to handle pool concurrency in *Golang*.

## Installation

```sh
go get github.com/francisbouvier/pool
```

## Example

Let's say you want to crawl a number of webpages in a concurrent way but limiting the number of simulteanous connections.

```go

import "github.com/francisbouvier/pool"

urls := []string{
	"https://www.google.com",
	"https://www.youtube.com",
	"https://www.slack.com",
	"https://www.twitter.com",
	"https://www.facebook.com",
	"https://www.docker.com",
	"https://www.snapchat.com",
}

concurrency := 2
p := pool.New(concurrency)
for _, url := range urls {
	p.Add() // Must be outside goroutines
	go func(url string) {
		// Your business logic
		_, err := http.Get(url)
		p.Done(err)
	}(url)
}
p.Wait()
p.Error() // to check errors
```
