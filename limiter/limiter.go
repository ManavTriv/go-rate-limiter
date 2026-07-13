package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu             sync.Mutex
	tokens         int
	maxTokens      int
	refillRate     time.Duration
	lastRefillTime time.Time
}