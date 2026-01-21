package http

import (
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/delivery/http/handlers"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/gin-gonic/gin"
	
)



func SetupRoutes(
	router *gin.Engine,
	authUseCase *usecase.AuthUseCase,
	tokenUseCase *usecase.TokenUseCase,
) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	authHandler := handlers.NewAuthHandler(authUseCase)
	tokenHandler := handlers.NewTokenHandler(tokenUseCase)
	passwordHandler := handlers.NewPasswordHandler(authUseCase)
	verificationHandler := handlers.NewVerificationHandler(authUseCase)
	permissionHandler := handlers.NewPermissionHandler(nil)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, utils.SuccessResponse(map[string]string{"status": "ok"}, "Service is healthy", utils.GenerateRequestID()))
	})

	public := router.Group("/auth")
	{
		public.POST("/register", authHandler.Register)
		public.POST("/login", authHandler.Login)
		public.POST("/logout", authHandler.Logout)
		public.POST("/refresh", tokenHandler.RefreshToken)
		
		public.POST("/password/reset/request", passwordHandler.RequestPasswordReset)
		public.POST("/password/reset", passwordHandler.ResetPassword)
		
		public.POST("/verification/request", verificationHandler.RequestEmailVerification)
		public.POST("/verification/confirm", verificationHandler.VerifyEmail)
	}

	protected := router.Group("/auth")
 
	protected.Use(middleware.ServiceAuthMiddleware())
	{
		protected.GET("/me", authHandler.GetMe)
		protected.POST("/validate", tokenHandler.ValidateToken)
		
		permissionsGroup := protected.Group("/permissions")
		{
			permissionsGroup.GET("/user", permissionHandler.GetUserPermissions)
			permissionsGroup.GET("/check", permissionHandler.CheckPermission)
		}
		
		adminGroup := protected.Group("/admin")
		adminGroup.Use(middleware.PermissionMiddleware(utils.PermReadUsers))
		{
			adminGroup.GET("/users", authHandler.GetAdminUsers)
		}
	}
}