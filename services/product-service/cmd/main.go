// Package main provides the product service for the e-commerce platform.
//
// @title Product Service API
// @version 1.0
// @description Product Service API for e-commerce platform
// @host localhost:8003
// @BasePath /product
// @schemes http
// @consumes application/json
// @produces application/json
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/logger"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/database"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init("admin-api-main")

	config := config.LoadConfig()

	// Construct the full DSN using individual config components
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.ProductDB,
		config.Database.Charset,
		config.Database.ParseTime,
		config.Database.Loc)

	db, err := database.InitDB(dsn)
	if err != nil {
		logger.Logger.Info("Error initializing database. Please check your database configuration.")
		os.Exit(1)
	}
	database.Migrate(db)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.RemoveExtraSlash = true
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.Cors.CorsOrigins, // Now correctly using a []string
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	err = r.SetTrustedProxies([]string{"0.0.0.0/0"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	//r.Static("/image_gallery", config.Storage.ImagePath)

	r.StaticFS("/image_gallery", gin.Dir(config.Storage.ImagePath, true))

	r.Use(gin.Recovery())
	routes.SetupRoutes(r, db)

	port := ":" + config.Services.ProductService.Port
	if err := r.Run(port); err != nil {
		logger.Logger.Fatalf("Error starting server: %v", err)
	}
}
