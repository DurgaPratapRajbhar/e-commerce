package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// SetupOrderRoutes configures order-related routes
func SetupOrderRoutes(api *gin.RouterGroup, sp *proxy.ServiceProxy, cfg *config.Config) {
	// Protected order routes
	protected := api.Group("/order")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		protected.Any("", sp.ProxyRequest("order"))
		protected.Any("/*path", sp.ProxyRequest("order"))
	}
}