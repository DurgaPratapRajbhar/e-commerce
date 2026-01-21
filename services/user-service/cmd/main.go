package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/logger"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/database"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Init("user-api-main")

	config := config.LoadConfig()

	fmt.Println("User Service is starting...",config.Database.UserDB)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.UserDB, config.Database.Charset, config.Database.ParseTime, config.Database.Loc)
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



	err = r.SetTrustedProxies([]string{"0.0.0.0/0"})
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	// if config.Storage.ImagePath != "" {
	// 	r.Static("/image_gallery", config.Storage.ImagePath)
	// }

	r.Use(gin.Recovery())
	routes.SetupRoutes(r, db)

	port := ":" + config.Services.UserService.Port
	if err := r.Run(port); err != nil {
		logger.Logger.Fatalf("Error starting server: %v", err)
	}
}
