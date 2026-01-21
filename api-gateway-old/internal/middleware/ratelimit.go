package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	limiters = make(map[string]*rate.Limiter)
	mu       sync.Mutex
)

// RateLimitMiddlewareWithConfig creates a rate limiter with custom config
func RateLimitMiddlewareWithConfig(requestsPerMinute int, burstSize int) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		mu.Lock()
		limiter, exists := limiters[clientIP]
		if !exists {
			// Create new limiter for this IP
			limiter = rate.NewLimiter(rate.Limit(float64(requestsPerMinute)/60.0), burstSize)
			limiters[clientIP] = limiter
		}
		mu.Unlock()

		// Check if request is allowed
		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "Rate limit exceeded. Please try again later.",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

// Cleanup old limiters periodically
func init() {
	go func() {
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()

		for range ticker.C {
			mu.Lock()
			// Clear all limiters every 5 minutes
			limiters = make(map[string]*rate.Limiter)
			mu.Unlock()
		}
	}()
}