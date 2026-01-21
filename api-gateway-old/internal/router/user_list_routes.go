package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/handler"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/gin-gonic/gin"
)

// SetupUserListRoutes configures user list-related routes
func SetupUserListRoutes(api *gin.RouterGroup, cfg *config.Config) {
	// Create clients
	authClient := client.NewAuthClient(cfg)
	userClient := client.NewUserClient(cfg)

	// Create handler
	userListHandler := handler.NewUserListHandler(authClient, userClient)

	// Protected user list routes
	protected := api.Group("/user-list")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		protected.Any("/all-users", userListHandler.UserList)
	}
}