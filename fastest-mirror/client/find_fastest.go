package client

import (
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/types"
	"net/http"
	"time"
)

// FindFastest receives a slice of URLs, sends a GET request to each URL concurrently, and returns an instance of
// types.Response for the URL that returns a response first.
func FindFastest(urls []string) types.Response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, urls := range urls {
		mirror := urls
		go func() {
			start := time.Now()
			_, err := http.Get(mirror + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirror
				latencyChan <- latency
			}
		}()
	}

	return types.Response{FastestURL: <-urlChan, Latency: <-latencyChan}
}
