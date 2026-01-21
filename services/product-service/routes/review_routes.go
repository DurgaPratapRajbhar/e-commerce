package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/controllers"
	repository "github.com/DurgaPratapRajbhar/e-commerce/product-service/repository/impl"
	services "github.com/DurgaPratapRajbhar/e-commerce/product-service/services/impl"
)

func ReviewRoutes(route *gin.RouterGroup, db *gorm.DB) {
	proReviewRepo := repository.NewProductReviewRepository(db)
	proReviewService := services.NewProductReviewService(proReviewRepo)
	proReviewController := controllers.NewProductReviewController(proReviewService)

	reviewGroup := route.Group("/product-reviews")
	reviewGroup.Use(middleware.ServiceAuthMiddleware())
	{
		reviewGroup.POST("", proReviewController.CreateReview)
		reviewGroup.GET("/:id", proReviewController.GetReview)
		reviewGroup.GET("", proReviewController.GetAllReviews)
		reviewGroup.PUT("/:id", proReviewController.UpdateReview)
		reviewGroup.DELETE("/:id", proReviewController.DeleteReview)
	}
}
