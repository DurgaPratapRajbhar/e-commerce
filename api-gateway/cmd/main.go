package main

import (
	 
	"log"
	"net/http"
	"time"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/router"
)

func main(){
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
   
	r := router.SetupRouter(cfg)

	server := &http.Server{
		Addr:         ":" + cfg.Server.Port,
		Handler:      r,
		ReadTimeout:  time.Duration(cfg.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.Server.WriteTimeout) * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	 
}