package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
)

type InventoryHandler struct {
	inventoryClient *client.InventoryClient
}

func NewInventoryHandler(inventoryClient *client.InventoryClient) *InventoryHandler {
	return &InventoryHandler{
		inventoryClient: inventoryClient,
	}
}

// GetInventory - Gin handler
func (h *InventoryHandler) GetInventory(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	inventory, err := h.inventoryClient.GetInventory(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    inventory,
	})
}