package gateway

import (
	"net/http"
	"sync"
	"time"
)

type TokenBucket struct {
	capacity      int
	currentTokens int
	rate          time.Duration
	lastRefill    time.Time
	mu            sync.Mutex
}

func NewTokenBucket(capacity int, rate time.Duration) *TokenBucket {
	return &TokenBucket{
		capacity:      capacity,
		currentTokens: capacity,
		rate:          rate,
		lastRefill:    time.Now(),
	}
}

func (t *TokenBucket) isAllowed() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	currentTime := time.Now()
	timePassedFromLastRefill := currentTime.Sub(t.lastRefill)

	newTokens := int(timePassedFromLastRefill / t.rate)
	if newTokens > 0 {
		t.currentTokens += newTokens
		if t.currentTokens > t.capacity {
			t.currentTokens = t.capacity
		}
		t.lastRefill = currentTime
	}

	if t.currentTokens > 0 {
		t.currentTokens--
		return true
	}
	return false
}

func rateLimitor(bucket *TokenBucket, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !bucket.isAllowed() {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
