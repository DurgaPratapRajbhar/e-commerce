package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// SetupProductRoutes configures product-related routes
func SetupProductRoutes(api *gin.RouterGroup, sp *proxy.ServiceProxy, cfg *config.Config) {
	// Protected product routes
	protected := api.Group("/product")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		protected.Any("", sp.ProxyRequest("product"))
		protected.Any("/*path", sp.ProxyRequest("product"))
	}
}