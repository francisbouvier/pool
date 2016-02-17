package pool

import (
	"net/http"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	concurrency := 2
	p := New(concurrency)
	urls := []string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.slack.com",
		"https://www.twitter.com",
		"https://www.facebook.com",
		"https://www.docker.com",
		"https://www.snapchat.com",
	}
	for _, url := range urls {
		p.Add()
		go func(url string) {
			defer p.Done()
			http.Get(url)
			time.Sleep(2)
			s := p.QueueSize()
			if s > concurrency {
				t.Errorf(
					"Max concurrency set to %d, but %d items in queue\n",
					concurrency, s,
				)
			}
		}(url)
	}
	p.Wait()
}
