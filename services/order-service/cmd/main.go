// Package main provides the order service for the e-commerce platform.
//
// @title Order Service API
// @version 1.0
// @description Order Service API for e-commerce platform
// @host localhost:8084
// @BasePath /api/v1
// @schemes http
// @consumes application/json
// @produces application/json
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/logger"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/database"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/routes"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init("admin-oredr")

	config := config.LoadConfig()

	// Construct the full DSN using individual config components
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.OrderDB,
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
		AllowOrigins: config.Cors.CorsOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	err = r.SetTrustedProxies([]string{"0.0.0.0/0"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	routes.SetupRoutes(r, db)
	port := ":" + config.Services.OrderService.Port

	if err := r.Run(port); err != nil {
		logger.Logger.Fatalf("Error starting server: %v", err)
	}
}
