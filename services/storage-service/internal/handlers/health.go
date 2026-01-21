package handlers

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/e-commerce/storage-service/internal/models"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HealthCheck(c *gin.Context) {
	response := models.HealthResponse{
		Status:  "healthy",
		Service: "storage-service",
		Version: "1.0.0",
	}

	c.JSON(http.StatusOK, response)
}

func (h *HealthHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "pong", "service": "storage-service"})
}