package pool

import (
	"net/http"
	"testing"
	"time"
)

var concurrency = 2

var urls = []string{
	"https://www.google.com",
	"https://www.youtube.com",
	"https://www.slack.com",
	"https://www.twitter.com",
	"https://www.facebook.com",
	"https://www.docker.com",
	"https://www.snapchat.com",
	"https://localhost:3000",
}

func TestPool(t *testing.T) {
	p := New(concurrency)
	for _, url := range urls {
		p.Add()
		go func(url string) {
			s := p.QueueSize()
			if s > concurrency {
				t.Errorf(
					"Max concurrency set to %d, but %d items in queue\n",
					concurrency, s,
				)
			}
			_, err := http.Get(url)
			time.Sleep(1)
			p.Done(err)
		}(url)
	}
	p.Wait()
	if p.Error() == nil {
		t.Error("At least 1 GET error")
	}
}
