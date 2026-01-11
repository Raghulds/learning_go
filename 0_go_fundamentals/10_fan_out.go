package gofundamentals

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var urls = []string{
	"https://go.dev",
	"https://ardanlabs.com",
	"https://ibm.com/no/such/page",
}

func FanoutDemo() {
	fmt.Println("Sequential calls")
	sequentialCalls()
	fmt.Println("--------------------------------")
	fmt.Println("Fanout calls for result")
	fanoutCalls_Result()
	fmt.Println("--------------------------------")
	fmt.Println("Fanout wait")
	fanoutCalls_Wait()
	fmt.Println("--------------------------------")
	fmt.Println("Fanout pool")
	fanoutCalls_Pool()
	fmt.Println("--------------------------------")
}

func fanoutCalls_Result() {
	start := time.Now()
	type Result struct {
		Url        string
		StatusCode int
		Error      error
	}

	var results []*Result
	ch := make(chan *Result)
	for _, v := range urls {
		go func() {
			r := &Result{
				Url: v,
			}
			defer func() {
				ch <- r
			}()
			r.StatusCode, r.Error = checkURL(v)
		}()
	}

	for range urls {
		res := <-ch
		fmt.Printf("Received: url %v, status %d, error %v \n", res.Url, res.StatusCode, res.Error)
		results = append(results, res)
	}

	duration := time.Since(start)
	fmt.Printf("%d urls in %v \n", len(urls), duration)
}

func fanoutCalls_Wait() {
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, v := range urls {
		go func() {
			defer func() {
				wg.Done()
			}()

			URLlog(v)
		}()
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Printf("%d urls in %v \n", len(urls), duration)

}

const PoolSize = 2

func fanoutCalls_Pool() {
	start := time.Now()
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(PoolSize)
	// Producers
	go func() {
		for _, v := range urls {
			ch <- v
		}
		close(ch)
	}()

	// Consumers
	for range PoolSize {
		go func() {
			defer wg.Done()
			for url := range ch {
				URLlog(url)
			}
		}()
	}
	wg.Wait()

	duration := time.Since(start)
	fmt.Printf("%d urls in %v \n", len(urls), duration)
}

func sequentialCalls() {

	start := time.Now()
	for _, v := range urls {
		stat, err := checkURL(v)
		fmt.Printf("%q: %d (%v)\n", v, stat, err)
	}
	duration := time.Since(start)
	fmt.Printf("%d urls in %v \n", len(urls), duration)
}

func checkURL(url string) (int, error) {
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	return res.StatusCode, nil
}

func URLlog(url string) {
	res, err := http.Get(url)
	if err != nil {
		slog.Error("urlLog", "url", url, "error", err)
		return
	}

	slog.Info("urlLog", "url", url, "status", res.StatusCode)
}
