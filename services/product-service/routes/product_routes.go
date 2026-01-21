package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/controllers"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository "github.com/DurgaPratapRajbhar/e-commerce/product-service/repository/impl"
	services "github.com/DurgaPratapRajbhar/e-commerce/product-service/services/impl"
)

func ProductRoutes(r *gin.RouterGroup, db *gorm.DB) {

	proRepo := repository.NewProductRepository(db)
	proService := services.NewProductService(proRepo)

	Controller := controllers.NewProductController(proService)

	userGroup := r.Group("/products")
	userGroup.Use(middleware.ServiceAuthMiddleware())
	{
		userGroup.POST("", Controller.CreateProduct)
		userGroup.GET("/:id", Controller.GetProduct)
		userGroup.GET("", Controller.GetAllProducts)
		userGroup.PUT("/:id", Controller.UpdateProduct)
		userGroup.DELETE("/:id", Controller.DeleteProduct)
	}
}
