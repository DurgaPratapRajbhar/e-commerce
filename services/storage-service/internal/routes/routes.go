package routes

import (
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/middleware"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/config"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/handlers"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/processor"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/storage"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/validator"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Initialize storage
	localStorage := storage.NewLocalStorage(cfg.Storage.UploadPath, cfg.Storage.BaseURL)

	// Initialize validator
	fileValidator := validator.NewFileValidator(
		cfg.Storage.MaxFileSize,
		cfg.Storage.AllowedTypes,
	)

	// Initialize processor
	imageProcessor := processor.NewImageProcessor(
		uint(cfg.Storage.ResizeWidth),
		uint(cfg.Storage.ResizeHeight),
		uint(cfg.Storage.ThumbnailWidth),
		uint(cfg.Storage.ThumbnailHeight),
	)

	// Initialize handlers
	uploadHandler := handlers.NewUploadHandler(localStorage, fileValidator, imageProcessor)
	serveHandler := handlers.NewServeHandler(localStorage)
	deleteHandler := handlers.NewDeleteHandler(localStorage)
	healthHandler := handlers.NewHealthHandler()

	// Health check (no auth)
	router.GET("/health", healthHandler.HealthCheck)
	router.GET("/ping", healthHandler.Ping)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Protected routes - requires authentication
		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware()) // Your global middleware from pkg
		{
			// Upload endpoints
			upload := protected.Group("/upload")
			{
				upload.POST("/product", uploadHandler.UploadProductImage)
				upload.POST("/avatar", uploadHandler.UploadAvatar)
				upload.POST("/document", uploadHandler.UploadDocument)
				upload.POST("/category-banner", uploadHandler.UploadCategoryBanner)
			}

			// File management (admin/authenticated users)
			files := protected.Group("/files")
			{
				files.DELETE("/:category/:subcategory/:filename", deleteHandler.DeleteFile)
				files.GET("/list", serveHandler.ListAllFiles)
				files.GET("/list/:category", serveHandler.ListFilesByCategory)
			}
		}
	}

	// Public static file serving (no auth) - for displaying images on frontend
	static := router.Group("/static")
	{
		// Product images
		static.GET("/products/images/full/:filename", serveHandler.ServeProductImageFull)
		static.GET("/products/images/medium/:filename", serveHandler.ServeProductImageMedium)
		static.GET("/products/images/thumbnails/:filename", serveHandler.ServeProductImageThumbnail)
		
		// User avatars
		static.GET("/users/avatars/:filename", serveHandler.ServeAvatar)
		
		// Category banners
		static.GET("/categories/banners/:filename", serveHandler.ServeCategoryBanner)
		
		// Documents (consider protecting these in production)
		static.GET("/documents/invoices/:filename", serveHandler.ServeInvoice)
		static.GET("/documents/receipts/:filename", serveHandler.ServeReceipt)
	}
}