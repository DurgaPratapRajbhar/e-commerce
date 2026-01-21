package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/handler"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/gin-gonic/gin"
)

// SetupUserRoutes configures user-related routes (simple direct approach)
func SetupUserRoutes(api *gin.RouterGroup, sp *proxy.ServiceProxy, cfg *config.Config) {
	// Create client
	userClient := client.NewUserClient(cfg)

	// Create handler
	userHandler := handler.NewUserHandler(userClient)

	// Protected user routes
	protected := api.Group("/user")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		protected.Any("/profiles", userHandler.GetUserProfiles)
	}
}