package handler

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/gin-gonic/gin"
)

type UserListHandler struct {
	authClient *client.AuthClient
	userClient *client.UserClient
}

func NewUserListHandler(authClient *client.AuthClient, userClient *client.UserClient) *UserListHandler {
	return &UserListHandler{
		authClient: authClient,
		userClient: userClient,
	}
}

// UserList fetches all users
func (h *UserListHandler) UserList(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Get users from user service
	users, err := h.userClient.GetUsers(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "user service unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
		"message": "Users fetched successfully",
	})
}