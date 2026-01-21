package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/handlers"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/repository/impl"
	serviceImpl "github.com/DurgaPratapRajbhar/e-commerce/services/order-service/services/impl"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	orderRepo := impl.NewOrderRepository(db)
	orderService := serviceImpl.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	orders := r.Group("/order")
	orders.Use(middleware.ServiceAuthMiddleware())
	{
		orders.POST("", orderHandler.CreateOrder)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.GET("/user/:userId", orderHandler.GetUserOrders)
		orders.GET("", orderHandler.GetAllOrders)
		orders.PUT("/:id", orderHandler.UpdateOrder)
		orders.DELETE("/:id", orderHandler.DeleteOrder)
		orders.PATCH("/:id/status", orderHandler.UpdateOrderStatus)
		orders.PATCH("/:id/payment", orderHandler.UpdatePaymentStatus)
	}
	
}