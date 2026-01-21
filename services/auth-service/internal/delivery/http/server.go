package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/config"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/usecase"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config       *config.Config
	authUseCase  *usecase.AuthUseCase
	tokenUseCase *usecase.TokenUseCase
}

func NewServer(cfg *config.Config, authUseCase *usecase.AuthUseCase, tokenUseCase *usecase.TokenUseCase) *Server {
	return &Server{
		config:       cfg,
		authUseCase:  authUseCase,
		tokenUseCase: tokenUseCase,
	}
}

func (s *Server) Start() error {
	router := gin.Default()
	
	SetupRoutes(router, s.authUseCase, s.tokenUseCase)
	
	
	srv := &http.Server{
		Addr:         ":" + s.config.Services.AuthService.Port,
		Handler:      router,
		ReadTimeout:  s.config.Server.ReadTimeout,
		WriteTimeout: s.config.Server.WriteTimeout,
	}
	
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	
	return srv.Shutdown(ctx)
}