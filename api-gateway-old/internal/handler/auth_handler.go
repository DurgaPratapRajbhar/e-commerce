package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
)

type AuthHandler struct {
	authClient *client.AuthClient
}

func NewAuthHandler(authClient *client.AuthClient) *AuthHandler {
	return &AuthHandler{
		authClient: authClient,
	}
}

// GetUsers - Gin handler
func (h *AuthHandler) GetUsers(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	users, err := h.authClient.GetUsers(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

// GetUserByID - Gin handler
func (h *AuthHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("id")
	authToken := c.GetHeader("Authorization")
	
	user, err := h.authClient.GetUserByID(c.Request.Context(), authToken, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

// CreateUser - Gin handler
func (h *AuthHandler) CreateUser(c *gin.Context) {
	var userData map[string]interface{}
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}
	
	authToken := c.GetHeader("Authorization")
	
	user, err := h.authClient.CreateUser(c.Request.Context(), authToken, userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    user,
	})
}

// UpdateUser - Gin handler
func (h *AuthHandler) UpdateUser(c *gin.Context) {
	userID := c.Param("id")
	authToken := c.GetHeader("Authorization")
	
	var userData map[string]interface{}
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}
	
	user, err := h.authClient.UpdateUser(c.Request.Context(), authToken, userID, userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

// DeleteUser - Gin handler
func (h *AuthHandler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")
	authToken := c.GetHeader("Authorization")
	
	err := h.authClient.DeleteUser(c.Request.Context(), authToken, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

// UpdateUserStatus - Gin handler
func (h *AuthHandler) UpdateUserStatus(c *gin.Context) {
	userID := c.Param("id")
	authToken := c.GetHeader("Authorization")
	
	var req struct {
		Status string `json:"status" binding:"required"`
	}
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid request body",
		})
		return
	}
	
	err := h.authClient.UpdateUserStatus(c.Request.Context(), authToken, userID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User status updated successfully",
	})
}