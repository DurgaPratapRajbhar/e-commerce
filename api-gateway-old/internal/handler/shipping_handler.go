package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
)

type ShippingHandler struct {
	shippingClient *client.ShippingClient
}

func NewShippingHandler(shippingClient *client.ShippingClient) *ShippingHandler {
	return &ShippingHandler{
		shippingClient: shippingClient,
	}
}

// GetShipments - Gin handler
func (h *ShippingHandler) GetShipments(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	shipments, err := h.shippingClient.GetShipments(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    shipments,
	})
}