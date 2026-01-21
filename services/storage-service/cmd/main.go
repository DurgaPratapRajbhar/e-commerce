package main

import (
	"log"
	"os"

	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/config"
	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Set Gin mode based on environment
	if cfg.AppEnv == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Create router
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Setup routes
	routes.SetupRoutes(router, cfg)

	// Start server
	port := os.Getenv("STORAGE_SERVICE_PORT")
	if port == "" {
		port = cfg.Server.Port
	}

	log.Printf("Storage service starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}