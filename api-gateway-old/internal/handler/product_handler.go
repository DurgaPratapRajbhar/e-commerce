package handler

import (
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productClient *client.ProductClient
}

func NewProductHandler(productClient *client.ProductClient) *ProductHandler {
	return &ProductHandler{
		productClient: productClient,
	}
}

// GetProducts - Single CRUD endpoint
func (h *ProductHandler) GetProducts(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	products, err := h.productClient.GetProducts(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "product service unavailable"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}