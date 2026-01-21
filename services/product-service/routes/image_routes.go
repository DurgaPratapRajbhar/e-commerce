package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/controllers"
	repository "github.com/DurgaPratapRajbhar/e-commerce/product-service/repository/impl"
	services "github.com/DurgaPratapRajbhar/e-commerce/product-service/services/impl"
)

func ImageRoutes(route *gin.RouterGroup, db *gorm.DB) {

	proRepo := repository.NewProductImagesRepository(db)
	proService := services.NewProductImagesService(proRepo)

	Controller := controllers.NewProductImageController(proService)
	imageGroup := route.Group("/product-images")
	imageGroup.Use(middleware.ServiceAuthMiddleware())
	{
		imageGroup.POST("", Controller.CreateImage)
		imageGroup.GET("/by-id/:id", Controller.GetImageByID)
		imageGroup.GET("/by-product/:product_id", Controller.GetAllImages)
		imageGroup.PUT("/:id", Controller.UpdateImage)
		imageGroup.DELETE("/:id", Controller.DeleteImage)
	}

}
