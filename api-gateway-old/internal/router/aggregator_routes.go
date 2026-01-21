package router

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/handler/aggregator"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/middleware"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/gin-gonic/gin"
	
	aggregator_service "github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/service/aggregator"
)

// SetupAggregatorRoutes configures aggregator-related routes (complex JOIN approach)
func SetupAggregatorRoutes(api *gin.RouterGroup, cfg *config.Config) {
	// Create clients
	sp := proxy.NewServiceProxy(cfg)
	aggProxy := proxy.NewAggregateProxy(sp)
	orderClient := client.NewOrderClient(cfg)
	userClient := client.NewUserClient(cfg)
	productClient := client.NewProductClient(cfg)
	cartClient := client.NewCartClient(cfg)

	// Create aggregator handlers
	orderAggregatorHandler := aggregator.NewOrderAggregatorHandler(orderClient, userClient, productClient, aggProxy)
	cartAggregator := aggregator_service.NewCartAggregator(cartClient, productClient, aggProxy)
	cartAggregatorHandler := aggregator.NewCartAggregatorHandler(cartAggregator)
	userAggregatorHandler := aggregator.NewUserAggregatorHandler(userClient, orderClient, productClient, cartClient, aggProxy)
	dashboardAggregatorHandler := aggregator.NewDashboardAggregatorHandler(orderClient, userClient, productClient, cartClient, aggProxy)

	// Protected aggregator routes
	protected := api.Group("/aggregator")
	protected.Use(middleware.AuthMiddleware(cfg.Auth.JWTSecret))
	{
		// Order aggregations
		protected.GET("/orders-with-users", orderAggregatorHandler.GetOrdersWithUsers)
		
		// Cart aggregations
		protected.GET("/cart-with-products", cartAggregatorHandler.GetCartWithProducts)
		
		// User aggregations
		protected.GET("/users/:id/profile-with-orders", userAggregatorHandler.GetUserProfileWithOrders)
		protected.GET("/users-with-order-count", userAggregatorHandler.GetUsersWithOrderCount)
		
		// Dashboard aggregations
		protected.GET("/dashboard", dashboardAggregatorHandler.GetDashboardData)
	}
}