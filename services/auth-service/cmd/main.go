package main

import (
	"log"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/delivery/http"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/infrastructure/database"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/infrastructure/database/mysql"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/infrastructure/token"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
)

func main() {
	cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }
  
	db := database.NewMySQLConnection(cfg.Database)
	defer db.Close()
	
	database.RunMigrations(db)
	
	userRepo := mysql.NewUserRepository(db)
	tokenService := token.NewJWTService(cfg.JWT.SecretKey, cfg.JWT.ExpiryHours)
	
	authUseCase := usecase.NewAuthUseCase(userRepo, tokenService)
	tokenUseCase := usecase.NewTokenUseCase(tokenService, userRepo)
	
	server := http.NewServer(cfg, authUseCase, tokenUseCase)
	
	log.Printf("Auth service starting on port %s", cfg.Services.AuthService.Port)
	log.Fatal(server.Start())
}