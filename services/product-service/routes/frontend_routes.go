package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository "github.com/DurgaPratapRajbhar/e-commerce/product-service/repository/impl"
	services "github.com/DurgaPratapRajbhar/e-commerce/product-service/services/impl"
)

func FrontendRoutes(r *gin.RouterGroup, db *gorm.DB) {

	prodRepo := repository.NewFrontendRepository(db)
	prodService := services.NewFrontendService(prodRepo)
	prodController := controllers.NewFrontendController(prodService)

	if prodController == nil {
		panic("prodController is nil in FrontendRoutes")
	}

	productGroup := r.Group("/frontend")
	{

		productGroup.GET("/category/:slug", prodController.GetProductsByCategorySlug)
		productGroup.GET("/product/:slug", prodController.GetProductData)
		productGroup.GET("/products", prodController.ProductSearch)

	}
}
