package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(requestNum int) {
			defer wg.Done()

			resp, err := http.Get("http://localhost:8080/")
			if err != nil {
				fmt.Printf("Request %d: error - %v\n", requestNum, err)
				return
			}
			defer resp.Body.Close()

			fmt.Printf("Request %d: status %d\n", requestNum, resp.StatusCode)
		}(i)
	}

	wg.Wait()
}