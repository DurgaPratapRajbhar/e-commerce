package handler

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userClient *client.UserClient
}

func NewUserHandler(userClient *client.UserClient) *UserHandler {
	return &UserHandler{
		userClient: userClient,
	}
}

// Direct handler: handler → service → client → microservice
// Simple, fast approach
func (h *UserHandler) GetUserProfiles(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	userIDsParam := c.Query("user_ids")

	if userIDsParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_ids parameter required"})
		return
	}

	userIDs := []string{} // In a real implementation, parse the user_ids from query param

	profiles, err := h.userClient.GetProfilesByUserIDs(c.Request.Context(), authToken, userIDs)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "user service unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}