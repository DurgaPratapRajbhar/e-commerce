package aggregator

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/service/aggregator"
	"github.com/gin-gonic/gin"
)

type OrderAggregatorHandler struct {
	orderAggregator *aggregator.OrderAggregator
}

func NewOrderAggregatorHandler(
	orderClient *client.OrderClient,
	userClient *client.UserClient,
	productClient *client.ProductClient,
	aggProxy *proxy.AggregateProxy,
) *OrderAggregatorHandler {
	orderAggregator := aggregator.NewOrderAggregator(orderClient, userClient, productClient, aggProxy)
	return &OrderAggregatorHandler{
		orderAggregator: orderAggregator,
	}
}

// JOIN Operations: handler/aggregator → service/aggregator → multiple clients → combine
// Complex, slower (but optimized with batching)
func (h *OrderAggregatorHandler) GetOrdersWithUsers(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Use aggregator service to fetch and combine data
	ordersWithUsers, err := h.orderAggregator.GetOrdersWithUsers(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ordersWithUsers,
		"message": "Orders with users fetched successfully",
	})
}