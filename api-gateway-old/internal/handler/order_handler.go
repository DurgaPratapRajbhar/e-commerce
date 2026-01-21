package handler

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderClient *client.OrderClient
}

func NewOrderHandler(orderClient *client.OrderClient) *OrderHandler {
	return &OrderHandler{
		orderClient: orderClient,
	}
}

// GetOrders - Single CRUD endpoint
func (h *OrderHandler) GetOrders(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	orders, err := h.orderClient.GetOrders(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "order service unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}