// Package main provides the payment service for the e-commerce platform.
//
// Schemes: http
// BasePath: /api/v1
// Version: 1.0.0
// Host: localhost:{{.Port}}
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"log"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/delivery/http"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/infrastructure/database"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/infrastructure/database/mysql"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/usecase"

)

func main() {
	// Initialize database connection
	db := database.NewMySQLConnection()
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get database connection: ", err)
	}
	defer sqlDB.Close()

	// Run migrations
	database.RunMigrations(db)

	// Initialize repositories
	paymentRepo := mysql.NewPaymentRepository(db)
	refundRepo := mysql.NewRefundRepository(db)

	// Initialize use cases
	paymentUseCase := usecase.NewPaymentUseCase(paymentRepo, refundRepo)

	// Initialize HTTP server
	server := http.NewServer(paymentUseCase)

	// The actual port is configured via environment variables and logged in the server.Start() method
	log.Fatal(server.Start())
}