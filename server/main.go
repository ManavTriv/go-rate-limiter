package main

import (
	"log"
	"net/http"
	"time"

	"go-rate-limiter/limiter"
)

func main() {
	bucket := limiter.NewTokenBucket(5, 200*time.Millisecond) // max 5 tokens and refills 1 every 200ms

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !bucket.Allow() { // no tokens left
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		w.Write([]byte("Request allowed\n")) // token consumed and continue
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}