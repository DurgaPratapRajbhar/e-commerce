package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// SetupShippingRoutes configures shipping-related routes
func SetupShippingRoutes(api *gin.RouterGroup, sp *proxy.ServiceProxy, cfg *config.Config) {
	// Protected shipping routes
	protected := api.Group("/shipment")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		protected.Any("", sp.ProxyRequest("shipment"))
		protected.Any("/*path", sp.ProxyRequest("shipment"))
	}
}