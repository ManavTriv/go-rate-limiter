package main

import (
	"log"
	"net/http"
	"time"

	"go-rate-limiter/limiter"
)

func main() {
	bucket := limiter.NewTokenBucket(5, 200*time.Millisecond)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !bucket.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		w.Write([]byte("Request allowed\n"))
	})

	log.Println("Server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}