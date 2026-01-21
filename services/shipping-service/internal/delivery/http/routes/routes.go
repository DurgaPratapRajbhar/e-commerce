package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/delivery/http/handlers"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, shippingUseCase *usecase.ShippingUseCase) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/shipment")
	api.Use(middleware.ServiceAuthMiddleware())
	{
		// Shipment CRUD operations
		shipments := api.Group("/")
		{
			// Static/specific routes first
			shipments.GET("all", handlers.GetShipments(shippingUseCase))
			shipments.GET("/order/:order_id", handlers.GetShipmentByOrderID(shippingUseCase))
			shipments.GET("/status/:status", handlers.GetShipmentsByStatus(shippingUseCase))
			
			// Parameterized routes last
			shipments.POST("", handlers.CreateShipment(shippingUseCase))
			shipments.GET("/:id", handlers.GetShipment(shippingUseCase))
			shipments.PUT("/:id", handlers.UpdateShipment(shippingUseCase))
			shipments.PATCH("/:id/status", handlers.UpdateShipmentStatus(shippingUseCase))
			shipments.DELETE("/:id", handlers.DeleteShipment(shippingUseCase))
		}

		// Tracking operations
		tracking := api.Group("/tracking")
		{
			tracking.POST("", handlers.CreateTrackingEvent(shippingUseCase))
			tracking.GET("/number/:tracking_number", handlers.GetShipmentByTrackingNumber(shippingUseCase))
			tracking.GET("/:shipment_id/latest", handlers.GetLatestTrackingEvent(shippingUseCase))
			tracking.GET("/:shipment_id", handlers.GetTrackingEvents(shippingUseCase))
		}
	}
}