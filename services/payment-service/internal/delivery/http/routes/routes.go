package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/delivery/http/handlers"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine, paymentUseCase *usecase.PaymentUseCase) {
	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Swagger docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
 
		payments := router.Group("/payment")
		payments.Use(middleware.ServiceAuthMiddleware())
		{
			payments.POST("", handlers.CreatePayment(paymentUseCase))
			payments.GET("/:id", handlers.GetPayment(paymentUseCase))
			payments.GET("/order/:order_id", handlers.GetPaymentByOrderID(paymentUseCase))
			payments.GET("/transaction/:transaction_id", handlers.GetPaymentByTransactionID(paymentUseCase))
			payments.GET("/user/:user_id", handlers.GetPaymentsByUserID(paymentUseCase))
			payments.PUT("/:id", handlers.UpdatePayment(paymentUseCase))
			payments.PATCH("/:id/status", handlers.UpdatePaymentStatus(paymentUseCase))
			payments.DELETE("/:id", handlers.DeletePayment(paymentUseCase))
			payments.GET("", handlers.GetPayments(paymentUseCase))
			payments.GET("/status/:status", handlers.GetPaymentsByStatus(paymentUseCase))
		}

		refunds := payments.Group("/refunds")
		{
			refunds.POST("", handlers.CreateRefund(paymentUseCase))
			refunds.GET("/:id", handlers.GetRefund(paymentUseCase))
			refunds.GET("/payment/:payment_id", handlers.GetRefundsByPaymentID(paymentUseCase))
			refunds.GET("/order/:order_id", handlers.GetRefundsByOrderID(paymentUseCase))
			refunds.GET("", handlers.GetRefunds(paymentUseCase))
			refunds.GET("/status/:status", handlers.GetRefundsByStatus(paymentUseCase))
		}
	
}