package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	// "io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
	"net/http/httptest"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
 
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

type ServiceProxy struct {
	config *config.Config
}

func NewServiceProxy(cfg *config.Config) *ServiceProxy {
	return &ServiceProxy{
		config: cfg,
	}
}

func (sp *ServiceProxy) ProxyRequest(serviceName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var serviceURL string
		var timeout time.Duration

		fmt.Printf("\n========== PROXY REQUEST ==========\n")
		fmt.Printf("Service: %s | Path: %s | Method: %s\n", serviceName, c.Request.URL.Path, c.Request.Method)

		// Get service URL and timeout
		switch serviceName {
		case "auth":
			serviceURL = sp.config.Services.AuthService.URL
			timeout = sp.config.Services.AuthService.Timeout
		case "user":
			serviceURL = sp.config.Services.UserService.URL
			timeout = sp.config.Services.UserService.Timeout
		case "product":
			serviceURL = sp.config.Services.ProductService.URL
			timeout = sp.config.Services.ProductService.Timeout
		case "cart":
			serviceURL = sp.config.Services.CartService.URL
			timeout = sp.config.Services.CartService.Timeout
		case "order":
			serviceURL = sp.config.Services.OrderService.URL
			timeout = sp.config.Services.OrderService.Timeout
		case "payment":
			serviceURL = sp.config.Services.PaymentService.URL
			timeout = sp.config.Services.PaymentService.Timeout
		case "shipment":
			serviceURL = sp.config.Services.ShippingService.URL
			timeout = sp.config.Services.ShippingService.Timeout
		case "inventory":
			serviceURL = sp.config.Services.InventoryService.URL
			timeout = sp.config.Services.InventoryService.Timeout
		default:
			c.JSON(http.StatusNotFound, utils.ErrorResponse(
				utils.ErrNotFound, "Service not found", nil, utils.GenerateRequestID(),
			))
			return
		}

		// Check if login/logout request
		originalPath := c.Request.URL.Path
		isLoginRequest := strings.HasSuffix(originalPath, "/login") && c.Request.Method == "POST"
		isLogoutRequest := strings.HasSuffix(originalPath, "/logout") && c.Request.Method == "POST"

		fmt.Printf("isLogin: %v | isLogout: %v\n", isLoginRequest, isLogoutRequest)

		// Parse target URL
		target, err := url.Parse(serviceURL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(
				utils.ErrInternalServer, "Invalid service URL", err.Error(), utils.GenerateRequestID(),
			))
			return
		}

		// Get user context
		userID := c.GetString("user_id")
		email := c.GetString("email")
		role := c.GetString("role")
		token := c.GetString("token")

		// Create reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(target)
		originalDirector := proxy.Director
		proxy.Director = func(req *http.Request) {
			originalDirector(req)
			
			// Strip /api/v1 prefix
			req.URL.Path = strings.TrimPrefix(originalPath, "/api/v1")
			req.URL.RawQuery = c.Request.URL.RawQuery
			req.Host = target.Host

			// Forward headers
			if userID != "" {
				req.Header.Set("X-User-ID", userID)
			}
			if email != "" {
				req.Header.Set("X-User-Email", email)
			}
			if role != "" {
				req.Header.Set("X-User-Role", role)
			}
			if token != "" {
				req.Header.Set("Authorization", "Bearer "+token)
			}

			// X-Forwarded headers
			req.Header.Set("X-Forwarded-For", c.ClientIP())
			req.Header.Set("X-Forwarded-Proto", c.Request.Proto)
			req.Header.Set("X-Forwarded-Host", c.Request.Host)

			fmt.Printf("Forwarding to: %s%s\n", target.String(), req.URL.Path)
		}

		// Error handler
		proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
			fmt.Printf("[ERROR] Service %s: %v\n", serviceName, err)
			c.JSON(http.StatusBadGateway, utils.ErrorResponse(
				utils.ErrExternalService, fmt.Sprintf("Service %s unavailable", serviceName),
				err.Error(), utils.GenerateRequestID(),
			))
		}

		// Set timeout
		if timeout > 0 {
			ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
			defer cancel()
			c.Request = c.Request.WithContext(ctx)
		}

		// Handle auth requests specially
		if isLoginRequest || isLogoutRequest {
			sp.handleAuthRequest(c, proxy, isLoginRequest, isLogoutRequest)
		} else {
			proxy.ServeHTTP(c.Writer, c.Request)
		}
	}
}

 
 func (sp *ServiceProxy) handleAuthRequest(c *gin.Context, proxy *httputil.ReverseProxy, isLogin, isLogout bool) {
	fmt.Println("[AUTH] Starting auth request handling...")
	
	// Create a response recorder to capture the backend response
	rec := httptest.NewRecorder()
	
	// Execute proxy with recorder
	proxy.ServeHTTP(rec, c.Request)
	
	// Get response details
	statusCode := rec.Code
	responseBody := rec.Body.Bytes()
	
	fmt.Printf("[AUTH] Response: Status=%d, BodyLen=%d\n", statusCode, len(responseBody))
	
	// Copy headers from recorder to actual response
	for k, v := range rec.Header() {
		c.Writer.Header()[k] = v
	}
	
	// Process response if successful
	if statusCode == http.StatusOK {
		if isLogin {
			sp.processLoginResponse(c, responseBody)
			return
		} else if isLogout {
			sp.processLogoutResponse(c, responseBody)
			return
		}
	}
	
	// Pass through for non-200 or non-auth requests
	c.Data(statusCode, rec.Header().Get("Content-Type"), responseBody)
}

func (sp *ServiceProxy) processLoginResponse(c *gin.Context, responseBody []byte) {
	fmt.Println("[LOGIN] Processing login response...")
	
	var authResponse map[string]interface{}
	if err := json.Unmarshal(responseBody, &authResponse); err != nil {
		fmt.Printf("[LOGIN] Parse error: %v\n", err)
		c.Data(http.StatusOK, "application/json", responseBody)
		return
	}

	// Extract token
	token, exists := authResponse["token"].(string)
	if !exists || token == "" {
		fmt.Println("[LOGIN] No token found in response")
		c.Data(http.StatusOK, "application/json", responseBody)
		return
	}

	fmt.Printf("[LOGIN] Token found (len=%d)\n", len(token))

	// Remove token from response
	// delete(authResponse, "token")

	// Set cookie
	maxAge := int(sp.config.Auth.TokenExpiry.Seconds())
	secure := sp.config.AppEnv == "production"
	
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   maxAge,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	}
	
	http.SetCookie(c.Writer, cookie)
	
	fmt.Printf("[LOGIN] Cookie set: %+v\n", cookie)
	
	// Send modified response (without token)
	c.JSON(http.StatusOK, authResponse)
	fmt.Println("[LOGIN] Response sent")
}

func (sp *ServiceProxy) processLogoutResponse(c *gin.Context, responseBody []byte) {
	fmt.Println("[LOGOUT] Processing logout...")
	
	// Clear cookie
	cookie := &http.Cookie{
		Name:     "auth_token",
		Value:    "",
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   sp.config.AppEnv == "production",
		SameSite: http.SameSiteLaxMode,
	}
	
	http.SetCookie(c.Writer, cookie)
	fmt.Println("[LOGOUT] Cookie cleared")

	// Send response
	c.Data(http.StatusOK, "application/json", responseBody)
}
// responseWriterWrapper captures response
type responseWriterWrapper struct {
	gin.ResponseWriter
	body       *bytes.Buffer
	statusCode int
}

func (w *responseWriterWrapper) Write(data []byte) (int, error) {
	w.body.Write(data)
	return w.ResponseWriter.Write(data)
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// HealthCheck checks all services
func (sp *ServiceProxy) HealthCheck(c *gin.Context) {
	healthStatus := make(map[string]string)

	services := map[string]string{
		"auth-service":      sp.config.Services.AuthService.URL,
		"user-service":      sp.config.Services.UserService.URL,
		"product-service":   sp.config.Services.ProductService.URL,
		"cart-service":      sp.config.Services.CartService.URL,
		"order-service":     sp.config.Services.OrderService.URL,
		"payment-service":   sp.config.Services.PaymentService.URL,
		"shipping-service":  sp.config.Services.ShippingService.URL,
		"inventory-service": sp.config.Services.InventoryService.URL,
	}

	for name, url := range services {
		if url != "" {
			status := checkServiceHealth(url + "/health")
			healthStatus[name] = status
		}
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(
		healthStatus, "Health check completed", utils.GenerateRequestID(),
	))
}

func checkServiceHealth(url string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "unhealthy"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "healthy"
	}
	return "unhealthy"
}