package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
)

type CartHandler struct {
	cartClient *client.CartClient
}

func NewCartHandler(cartClient *client.CartClient) *CartHandler {
	return &CartHandler{
		cartClient: cartClient,
	}
}

// GetCartItems - Gin handler
func (h *CartHandler) GetCartItems(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	items, err := h.cartClient.GetCartItems(c.Request.Context(), authToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
	})
}