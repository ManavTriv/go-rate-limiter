package limiter

import (
	"sync"
	"time"
)

type TokenBucket struct {
	mu sync.Mutex // guards tokens and lastRefillTime from concurrent access
	tokens int
	maxTokens int
	refillRate time.Duration // how often 1 token is added
	lastRefillTime time.Time
}

func NewTokenBucket(maxTokens int, refillRate time.Duration) *TokenBucket {
	return &TokenBucket{
		tokens: maxTokens, // start full
		maxTokens: maxTokens,
		refillRate: refillRate,
		lastRefillTime: time.Now(),
	}
}

func (tb *TokenBucket) Allow() bool {
	tb.mu.Lock()
	defer tb.mu.Unlock()

	tb.refill() // top up before checking

	if tb.tokens > 0 {
		tb.tokens--
		return true
	}
	return false
}

func (tb *TokenBucket) refill() {
	// called from within Allow(), which already holds the lock 
	now := time.Now()
	elapsed := now.Sub(tb.lastRefillTime)

	tokensToAdd := int(elapsed / tb.refillRate) // whole tokens earned since last refill

	if tokensToAdd > 0 {
		tb.tokens += tokensToAdd
		if tb.tokens > tb.maxTokens {
			tb.tokens = tb.maxTokens // cap at bucket capacity
		}
		tb.lastRefillTime = now // only advance the clock when tokens were actually added
	}
}