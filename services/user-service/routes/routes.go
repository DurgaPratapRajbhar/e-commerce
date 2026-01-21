package routes

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/controllers"
	repository "github.com/DurgaPratapRajbhar/e-commerce/user-service/repository/impl"
	service "github.com/DurgaPratapRajbhar/e-commerce/user-service/services/impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(route *gin.Engine, db *gorm.DB) {
	route.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	userProfileRepo := repository.NewUserProfileRepository(db)
	userAddressRepo := repository.NewUserAddressRepository(db)

	userProfileService := service.NewUserProfileService(userProfileRepo)
	userAddressService := service.NewUserAddressService(userAddressRepo)

	userProfileController := controllers.NewUserProfileController(userProfileService)
	userAddressController := controllers.NewUserAddressController(userAddressService)

	userRoutes := route.Group("/user")
	
	userRoutes.Use(middleware.ServiceAuthMiddleware())
	
	userRoutes.Use(extractUserFromHeaders())

	profileGroup := userRoutes.Group("/profiles")
	{
	 
		profileGroup.POST("", userProfileController.CreateProfile)
		profileGroup.GET("/:user_id", userProfileController.GetProfileByUserID)
		profileGroup.GET("/id/:id", userProfileController.GetProfileByID)
		profileGroup.GET("/bulk", userProfileController.GetProfilesBulk)
		profileGroup.PUT("/:user_id", userProfileController.UpdateProfile)
		profileGroup.DELETE(":user_id", userProfileController.DeleteProfile)
	}

	addressGroup := userRoutes.Group("/addresses")
	{
		addressGroup.POST("", userAddressController.CreateAddress)
		addressGroup.GET("/:id", userAddressController.GetAddressByID)
		addressGroup.GET("/user/:user_id", userAddressController.GetAddressesByUserID)
		addressGroup.GET("/user/:user_id/type/:type", userAddressController.GetAddressByUserIDAndType)
		addressGroup.PUT("/:id", userAddressController.UpdateAddress)
		addressGroup.DELETE("/:id", userAddressController.DeleteAddress)
		addressGroup.PUT("/user/:user_id/default/:address_id", userAddressController.SetDefaultAddress)
		addressGroup.GET("/user/:user_id/default", userAddressController.GetDefaultAddress)
	}
}

func extractUserFromHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetHeader("X-User-ID")
		email := c.GetHeader("X-User-Email")
		role := c.GetHeader("X-User-Role")

		if userID != "" {
			c.Set("user_id", userID)
		}
		if email != "" {
			c.Set("email", email)
		}
		if role != "" {
			c.Set("role", role)
		}

		c.Next()
	}
}