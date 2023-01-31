package client

import (
	"github.com/vltraheaven/go-microservice-experiments/fastest-mirror/types"
	"net/http"
	"time"
)

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
