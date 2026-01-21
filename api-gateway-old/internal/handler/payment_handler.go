package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
)

type PaymentHandler struct {
	paymentClient *client.PaymentClient
}

func NewPaymentHandler(paymentClient *client.PaymentClient) *PaymentHandler {
	return &PaymentHandler{
		paymentClient: paymentClient,
	}
}

// GetPayments - Gin handler
func (h *PaymentHandler) GetPayments(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	payments, err := h.paymentClient.GetPayments(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    payments,
	})
}