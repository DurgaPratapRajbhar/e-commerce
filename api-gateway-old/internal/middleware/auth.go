package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthMiddleware validates JWT token from cookie and sets user context
// func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		fmt.Println("[AUTH-MW] Checking authentication...")

// 		var tokenString string

// 		// Try cookie first
// 		cookie, err := c.Cookie("auth_token")
// 		if err == nil && cookie != "" {
// 			tokenString = cookie
// 			fmt.Printf("[AUTH-MW] Token from cookie (len=%d)\n", len(cookie))
// 		} else {
// 			// Try Authorization header as fallback
// 			authHeader := c.GetHeader("Authorization")
// 			if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
// 				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
// 				fmt.Printf("[AUTH-MW] Token from Authorization header (len=%d)\n", len(tokenString))
// 			}
// 		}

// 		if tokenString == "" {
// 			fmt.Printf("[AUTH-MW] No token found - Cookie error: %v, Auth header: %s\n", err, c.GetHeader("Authorization"))
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Unauthorized - No authentication token",
// 			})
// 			c.Abort()
// 			return
// 		}

// 		fmt.Printf("[AUTH-MW] JWT Secret being used (len=%d): %s...\n", len(jwtSecret), jwtSecret[:min(10, len(jwtSecret))])

// 		// Parse and validate JWT
// 		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 			// Validate signing method
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 			}
// 			return []byte(jwtSecret), nil
// 		})

// 		if err != nil {
// 			fmt.Printf("[AUTH-MW] Token parse error: %v\n", err)
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid authentication token",
// 				"details": err.Error(),
// 			})
// 			c.Abort()
// 			return
// 		}

// 		// Extract claims
// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			userID := fmt.Sprintf("%.0f", claims["user_id"].(float64))
// 			email := claims["email"].(string)
// 			role := claims["role"].(string)

// 			fmt.Printf("[AUTH-MW] Authenticated: user_id=%s, email=%s, role=%s\n", userID, email, role)

// 			// Set in context for proxy to use
// 			c.Set("user_id", userID)
// 			c.Set("email", email)
// 			c.Set("role", role)
// 			c.Set("token", tokenString) // Pass original token

// 			c.Next()
// 		} else {
// 			fmt.Println("[AUTH-MW] Invalid token claims")
// 			c.JSON(http.StatusUnauthorized, gin.H{
// 				"error": "Invalid token claims",
// 			})
// 			c.Abort()
// 			return
// 		}
// 	}
// }

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[AUTH-MW] Checking authentication...")

		var tokenString string

		// OPTION 1: Check Authorization header (Mobile, Postman, APIs)
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			fmt.Printf("[AUTH-MW] Token from Authorization header (len=%d)\n", len(tokenString))
		} else {
			// OPTION 2: Check Cookie (Browser)
			cookie, err := c.Cookie("auth_token")
			if err == nil && cookie != "" {
				tokenString = cookie
				fmt.Printf("[AUTH-MW] Token from cookie (len=%d)\n", len(cookie))
			}
		}

		if tokenString == "" {
			fmt.Println("[AUTH-MW] No token found")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized - No authentication token",
			})
			c.Abort()
			return
		}

		// JWT validation (same code as before)
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid authentication token",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID := fmt.Sprintf("%.0f", claims["user_id"].(float64))
			email := claims["email"].(string)
			role := claims["role"].(string)

			// Set context for proxy to forward
			c.Set("user_id", userID)
			c.Set("email", email)
			c.Set("role", role)
			c.Set("token", tokenString)

			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
		}
	}
}


// Gateway - Combined Auth Middleware
// func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		// 1. Check for Service API Key (Service-to-Service)
// 		if apiKey := c.GetHeader("X-API-Key"); apiKey != "" {
// 			if validateServiceAPIKey(apiKey) {
// 				c.Set("is_service", true)
// 				c.Set("service_name", getServiceName(apiKey))
// 				c.Next()
// 				return
// 			}
// 		}
		
// 		// 2. Check for User Token (Cookie or Header)
// 		var token string
		
// 		// Try Authorization header (Mobile/API)
// 		if auth := c.GetHeader("Authorization"); strings.HasPrefix(auth, "Bearer ") {
// 			token = strings.TrimPrefix(auth, "Bearer ")
// 		} else {
// 			// Try Cookie (Browser)
// 			token, _ = c.Cookie("auth_token")
// 		}
		
// 		if token == "" {
// 			c.JSON(401, gin.H{"error": "No authentication"})
// 			c.Abort()
// 			return
// 		}
		
// 		// Validate JWT
// 		claims, err := validateJWT(token, cfg.Auth.JWTSecret)
// 		if err != nil {
// 			c.JSON(401, gin.H{"error": "Invalid token"})
// 			c.Abort()
// 			return
// 		}
		
// 		// Set user context
// 		c.Set("user_id", claims.UserID)
// 		c.Set("email", claims.Email)
// 		c.Set("role", claims.Role)
// 		c.Next()
// 	}
// }

// OptionalAuthMiddleware allows both authenticated and public access
func OptionalAuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie("auth_token")
		if err == nil && cookie != "" {
			token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(jwtSecret), nil
			})

			if err == nil {
				if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
					userID := fmt.Sprintf("%.0f", claims["user_id"].(float64))
					c.Set("user_id", userID)
					c.Set("email", claims["email"].(string))
					c.Set("role", claims["role"].(string))
					c.Set("token", cookie)
				}
			}
		}
		c.Next()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}