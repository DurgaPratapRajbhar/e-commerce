package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/delivery/http/handlers"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
)

// func SetupRoutes(router *gin.Engine, inventoryUseCase *usecase.InventoryUseCase) {
// 	// Health check endpoint
// 	router.GET("/health", func(c *gin.Context) {
// 		c.JSON(200, gin.H{"status": "ok"})
// 	})

// 	// Swagger docs
// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

// 	// API routes
// 	// api := router.Group("")
// 	// {
// 	// 	inventory := api.Group("/inventory")
// 	// 	{
// 	// 		inventory.POST("", handlers.CreateInventory(inventoryUseCase))
// 	// 		inventory.GET("/product/:product_id/variant/:variant_id", handlers.GetInventory(inventoryUseCase))
// 	// 		inventory.PUT("/update", handlers.UpdateInventory(inventoryUseCase))
// 	// 		inventory.GET("/low-stock", handlers.GetLowStockItems(inventoryUseCase))
// 	// 		inventory.DELETE("/product/:product_id/variant/:variant_id", handlers.DeleteInventory(inventoryUseCase))
// 	// 	}

// 	// 	transactions := api.Group("/inventory/transactions")
// 	// 	{
// 	// 		transactions.POST("", handlers.CreateTransaction(inventoryUseCase))
// 	// 		transactions.GET("/product/:product_id", handlers.GetTransactionsByProduct(inventoryUseCase))
// 	// 		transactions.GET("/product/:product_id/variant/:variant_id", handlers.GetTransactionsByProductAndVariant(inventoryUseCase))
// 	// 		transactions.GET("/reference/:reference_id", handlers.GetTransactionsByReference(inventoryUseCase))
// 	// 		transactions.GET("/recent", handlers.GetRecentTransactions(inventoryUseCase))
// 	// 	}
// 	// }

// 	api := router.Group("inventory")
// 	{
// 		inventory := api.Group("/inventory")
// 		{
// 			inventory.POST("", handlers.CreateInventory(inventoryUseCase))
// 			inventory.GET("/product/:product_id/variant/:variant_id", handlers.GetInventory(inventoryUseCase))
// 			inventory.PUT("/product/:product_id/variant/:variant_id", handlers.UpdateInventory(inventoryUseCase))
// 			inventory.GET("/low-stock", handlers.GetLowStockItems(inventoryUseCase))
// 			inventory.DELETE("/product/:product_id/variant/:variant_id", handlers.DeleteInventory(inventoryUseCase))
// 		}

// 		transactions := api.Group("/transactions")
// 		{
// 			transactions.POST("", handlers.CreateTransaction(inventoryUseCase))
// 			transactions.GET("/product/:product_id", handlers.GetTransactionsByProduct(inventoryUseCase))
// 			transactions.GET("/product/:product_id/variant/:variant_id", handlers.GetTransactionsByProductAndVariant(inventoryUseCase))
// 			transactions.GET("/reference/:reference_id", handlers.GetTransactionsByReference(inventoryUseCase))
// 			transactions.GET("/recent", handlers.GetRecentTransactions(inventoryUseCase))
// 		}
// 	}
// }

func SetupRoutes(router *gin.Engine, inventoryUseCase *usecase.InventoryUseCase) {

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, utils.SuccessResponse(map[string]string{"status": "ok"}, "Service is healthy", utils.GenerateRequestID()))
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	inventory := router.Group("/inventory")
	inventory.Use(middleware.ServiceAuthMiddleware())
	{
		inventory.POST("", handlers.CreateInventory(inventoryUseCase))
		inventory.GET("", handlers.GetInventoryList(inventoryUseCase)) // Add this route to handle GET /inventory
		inventory.GET("/product/:product_id/variant/:variant_id", handlers.GetInventory(inventoryUseCase))
		inventory.PUT("", handlers.UpdateInventory(inventoryUseCase))
		inventory.GET("/low-stock", handlers.GetLowStockItems(inventoryUseCase))
		inventory.DELETE("/product/:product_id/variant/:variant_id", handlers.DeleteInventory(inventoryUseCase))

		transactions := inventory.Group("/transactions")
		{
			transactions.POST("", handlers.CreateTransaction(inventoryUseCase))
			transactions.GET("/product/:product_id", handlers.GetTransactionsByProduct(inventoryUseCase))
			transactions.GET("/product/:product_id/variant/:variant_id", handlers.GetTransactionsByProductAndVariant(inventoryUseCase))
			transactions.GET("/reference/:reference_id", handlers.GetTransactionsByReference(inventoryUseCase))
			transactions.GET("/recent", handlers.GetRecentTransactions(inventoryUseCase))
		}
	}
}
