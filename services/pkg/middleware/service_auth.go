package middleware

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ServiceAuthMiddleware validates headers from API Gateway
// Used by ALL microservices (not API Gateway)
func ServiceAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("\n========== SERVICE AUTH MIDDLEWARE ==========")
		fmt.Printf("Path: %s | Method: %s\n", c.Request.URL.Path, c.Request.Method)
		
		// Debug: Print ALL headers
		fmt.Println("All Headers:")
		for name, values := range c.Request.Header {
			for _, value := range values {
				fmt.Printf("  %s: %s\n", name, value)
			}
		}
		
		// Check if user info is provided in headers from API Gateway
		userIDStr := c.GetHeader("X-User-ID")
		email := c.GetHeader("X-User-Email")
		role := c.GetHeader("X-User-Role")

		fmt.Printf("X-User-ID: '%s'\n", userIDStr)
		fmt.Printf("X-User-Email: '%s'\n", email)
		fmt.Printf("X-User-Role: '%s'\n", role)

		if userIDStr == "" {
			fmt.Println("[ERROR] Missing X-User-ID header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - Missing user information"})
			c.Abort()
			return
		}

		userID, err := strconv.ParseUint(userIDStr, 10, 64)
		if err != nil {
			fmt.Printf("[ERROR] Invalid user ID format: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID format"})
			c.Abort()
			return
		}

		// Set in context for handlers
		c.Set("user_id", uint(userID))
		c.Set("email", email)
		c.Set("role", role)

		fmt.Printf("[SUCCESS] User authenticated: ID=%d, Email=%s, Role=%s\n", userID, email, role)
		fmt.Println("============================================\n")

		c.Next()
	}
}

// OptionalServiceAuth allows both authenticated and public access
func OptionalServiceAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		userIDStr := c.GetHeader("X-User-ID")
		if userIDStr != "" {
			userID, err := strconv.ParseUint(userIDStr, 10, 64)
			if err == nil {
				c.Set("user_id", uint(userID))
				c.Set("email", c.GetHeader("X-User-Email"))
				c.Set("role", c.GetHeader("X-User-Role"))
				fmt.Printf("[OPTIONAL-AUTH] User context set: ID=%d\n", userID)
			}
		} else {
			fmt.Println("[OPTIONAL-AUTH] No user headers found, proceeding as guest")
		}
		c.Next()
	}
}

// RequireRole checks if user has required role
func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied - No role information"})
			c.Abort()
			return
		}

		userRole := role.(string)
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied - Insufficient permissions"})
		c.Abort()
	}
}