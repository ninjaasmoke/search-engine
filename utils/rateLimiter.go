package utils

import (
	"net/http"
	"sync"
	"time"
)

// RateLimiter defines a rate limiter structure
type RateLimiter struct {
	mu         sync.Mutex
	visitors   map[string]*visitor
	maxPerIP   int
	expiration time.Duration
}

type visitor struct {
	seen   time.Time
	visits int
}

// NewRateLimiter creates a new instance of RateLimiter
func NewRateLimiter(maxPerIP int, expiration time.Duration) *RateLimiter {
	limiter := &RateLimiter{
		visitors:   make(map[string]*visitor),
		maxPerIP:   maxPerIP,
		expiration: expiration,
	}
	go limiter.cleanupExpiredVisitors()
	return limiter
}

// cleanupExpiredVisitors is a goroutine to clean up expired visitors from the map
func (limiter *RateLimiter) cleanupExpiredVisitors() {
	for {
		time.Sleep(limiter.expiration)
		limiter.mu.Lock()
		for ip, v := range limiter.visitors {
			if time.Since(v.seen) > limiter.expiration {
				delete(limiter.visitors, ip)
			}
		}
		limiter.mu.Unlock()
	}
}

// Limit implements the rate limiting logic
func (limiter *RateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr // Get IP address of the client

		limiter.mu.Lock()
		v, exists := limiter.visitors[ip]
		if !exists {
			v = &visitor{seen: time.Now()}
			limiter.visitors[ip] = v
		}
		v.visits++
		limiter.mu.Unlock()

		// Check if the number of visits exceeds the limit
		if v.visits > limiter.maxPerIP {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
