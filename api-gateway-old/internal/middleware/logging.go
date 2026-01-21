package middleware

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
)

// LoggingMiddleware provides custom request logging
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		// Log request details
		fmt.Printf("[API-GATEWAY] %s | %3d | %13v | %15s | %-7s %s\n",
			startTime.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}