package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup // counter for how many goroutines are still running

	for i := 1; i <= 10; i++ {
		wg.Add(1) // one more goroutine to wait for

		go func(requestNum int) { // launches as go rountine which runs concurrently
			defer wg.Done() // signal this goroutine is done when it returns

			resp, err := http.Get("http://localhost:8080/")
			if err != nil {
				fmt.Printf("Request %d: error - %v\n", requestNum, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Request %d: status %d\n", requestNum, resp.StatusCode)
		}(i) // pass i in explicitly to avoid all goroutines sharing the same loop variable
	}

	wg.Wait() // wait until all goroutines have called Done()
}