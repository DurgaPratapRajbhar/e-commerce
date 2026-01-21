package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/database"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	config := config.LoadConfig()
	fmt.Printf("Cart service port: %s\n", config.Services.CartService.Port)
	fmt.Printf("Cart DB DSN: %s\n", config.Database.CartDB)

	database.ConnectDB(config)

	err = database.DB.AutoMigrate(&models.Cart{})
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}

	fmt.Println("Database migrated successfully.")

	router := gin.Default()

	routes.SetupCartRoutes(router)

	port := config.Services.CartService.Port
	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8083" // default cart service port
		}
	}
	log.Printf("Cart service is running on port %s", ":" + port)
	router.Run(":" + port)
}