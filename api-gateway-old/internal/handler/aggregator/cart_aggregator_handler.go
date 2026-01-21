package aggregator

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/service/aggregator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartAggregatorHandler struct {
	cartAggregator *aggregator.CartAggregator
}

func NewCartAggregatorHandler(cartAggregator *aggregator.CartAggregator) *CartAggregatorHandler {
	return &CartAggregatorHandler{
		cartAggregator: cartAggregator,
	}
}

// JOIN Endpoint: Cart with products
func (h *CartAggregatorHandler) GetCartWithProducts(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	// Using the aggregator to join cart and product data
	cartWithProducts, err := h.cartAggregator.GetCartWithProducts(c.Request.Context(), authToken)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "aggregation service unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cart_with_products": cartWithProducts,
	})
}