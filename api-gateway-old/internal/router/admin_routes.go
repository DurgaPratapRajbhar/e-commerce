package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
)

// SetupAdminRoutes configures admin-related routes
func SetupAdminRoutes(api *gin.RouterGroup, sp *proxy.ServiceProxy, ap *proxy.AggregateProxy, cfg *config.Config) {
	// Admin routes (protected)
	admin := api.Group("/admin")
	admin.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		admin.Any("/users", ap.AggregateAdminUsers)
	}
}