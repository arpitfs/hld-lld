package gateway

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func setUpRateLimiter(router *mux.Router) http.Handler {
	tokenBucket := NewTokenBucket(BucketCapacity, BucketRefillingRate*time.Second)
	rateLimitMiddleware := rateLimitor(tokenBucket, router)

	return rateLimitMiddleware
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
