package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/delivery/http/routes"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/gin-gonic/gin"
)

type Server struct {
	paymentUseCase *usecase.PaymentUseCase
	router         *gin.Engine
}

func NewServer(paymentUseCase *usecase.PaymentUseCase) *Server {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	s := &Server{
		paymentUseCase: paymentUseCase,
		router:         router,
	}

	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	routes.SetupRoutes(s.router, s.paymentUseCase)
}

func (s *Server) Start() error {
	// Load the shared config
	globalConfig := config.LoadConfig()
	server := &http.Server{
		Addr:    ":" + globalConfig.Services.PaymentService.Port,
		Handler: s.router,
	}

	log.Printf("Payment service starting on port %s", globalConfig.Services.PaymentService.Port)

	// Graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	log.Println("Server exited")
	return nil
}