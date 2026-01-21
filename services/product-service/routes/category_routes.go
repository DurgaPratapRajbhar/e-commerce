package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/controllers"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	repository "github.com/DurgaPratapRajbhar/e-commerce/product-service/repository/impl"
	services "github.com/DurgaPratapRajbhar/e-commerce/product-service/services/impl"
)

func CategoryRoutes(r *gin.RouterGroup, db *gorm.DB) {

	catRepo := repository.NewCategoryRepository(db)
	catService := services.NewCategoryService(catRepo)

	Controller := controllers.NewCategoryController(catService)
	userGroup := r.Group("/categories")
	userGroup.Use(middleware.ServiceAuthMiddleware())
	{
		userGroup.POST("", Controller.CreateCategory)
		userGroup.GET("/:id", Controller.GetCategory)
		userGroup.GET("", Controller.GetAllCategories)
		userGroup.PUT("/:id", Controller.UpdateCategory)
		userGroup.DELETE("/:id", Controller.DeleteCategory)
		userGroup.GET("/list", Controller.GetAllCategoriesList)

	}
}
