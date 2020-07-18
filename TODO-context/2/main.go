package main

import (
	"context"
	"net/http"
	"time"
)

// https://xnum.github.io/2018/06/golang-context

func fetcher(ctx context.Context, url string) {
	req, err := http.NewRequest(http.Get, url, nil)
	req = req.WithContext(ctx)
	client := &http.Client{}
	client.Do(req)
}

func fetcher(pctx context.Context, url string) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	req, err := http.NewRequest(http.Get, url, nil)
	req = req.WithContext(ctx)
	client := &http.Client{}
	client.Do(req)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// Now I have a derived context to control my goroutines.
	go fetcher(ctx, "http://www.google.com")
	go fetcher(ctx, "http://www.google.com/1")
	go fetcher(ctx, "http://www.google.com/2")
	go fetcher(ctx, "http://www.google.com/3")

	doSomething()

	// I don't want to wait anymore...
	cancel()
}
