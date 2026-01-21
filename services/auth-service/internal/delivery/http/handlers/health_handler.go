package handlers

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	data := map[string]string{
		"status":  "ok",
		"message": "auth-service is running",
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Service is healthy", requestID))
}

func ReadinessCheck(c *gin.Context) {
	data := map[string]string{
		"status":  "ok",
		"message": "auth-service is ready",
	}
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(data, "Service is healthy", requestID))
}