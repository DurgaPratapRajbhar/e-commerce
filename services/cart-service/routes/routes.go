package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/controllers"
	repositoryImpl "github.com/DurgaPratapRajbhar/e-commerce/cart-service/repository/impl"
	serviceImpl "github.com/DurgaPratapRajbhar/e-commerce/cart-service/services/impl"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func SetupCartRoutes(router *gin.Engine) {
	 
	cartRepo := repositoryImpl.NewCartRepository()
	cartService := serviceImpl.NewCartService(cartRepo)
	cartController := controllers.NewCartController(cartService)

	cartRoutes := router.Group("cart")
	cartRoutes.Use(middleware.ServiceAuthMiddleware())
	{
		cartRoutes.POST("", cartController.AddToCart)
		cartRoutes.GET("/:id", cartController.GetCartByID)
		cartRoutes.GET("/user/:userId", cartController.GetCartByUserID)
		cartRoutes.PUT("/:id", cartController.UpdateCart)
		cartRoutes.DELETE("/:id", cartController.RemoveFromCart)
		cartRoutes.DELETE("/user/:userId", cartController.ClearCart)
	}
}