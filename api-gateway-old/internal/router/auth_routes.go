package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// SetupAuthRoutes configures authentication-related routes
func SetupAuthRoutes(api *gin.RouterGroup, sp *proxy.ServiceProxy, cfg *config.Config) {
	// Auth routes (public)
	auth := api.Group("/auth")
	{
		auth.Any("/register", sp.ProxyRequest("auth"))
		auth.Any("/login", sp.ProxyRequest("auth"))
		auth.Any("/logout", sp.ProxyRequest("auth"))
		auth.Any("/refresh", sp.ProxyRequest("auth"))
		auth.Any("/password/*path", sp.ProxyRequest("auth"))
		auth.Any("/verification/*path", sp.ProxyRequest("auth"))

		// Protected auth routes
		authProtected := auth.Group("")
		authProtected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
		{
			authProtected.Any("/me", sp.ProxyRequest("auth"))
			authProtected.Any("/validate", sp.ProxyRequest("auth"))
			authProtected.Any("/permissions/*path", sp.ProxyRequest("auth"))
		}
	}
}