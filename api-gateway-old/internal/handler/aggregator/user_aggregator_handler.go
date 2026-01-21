package aggregator

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/service/aggregator"
	"github.com/gin-gonic/gin"
)

type UserAggregatorHandler struct {
	userAggregator *aggregator.UserAggregator
}

func NewUserAggregatorHandler(
	userClient *client.UserClient,
	orderClient *client.OrderClient,
	productClient *client.ProductClient,
	cartClient *client.CartClient,
	aggProxy *proxy.AggregateProxy,
) *UserAggregatorHandler {
	userAggregator := aggregator.NewUserAggregator(userClient, orderClient, productClient, cartClient, aggProxy)
	return &UserAggregatorHandler{
		userAggregator: userAggregator,
	}
}

// JOIN Operations: handler/aggregator → service/aggregator → multiple clients → combine
// Complex, slower (but optimized with batching)
func (h *UserAggregatorHandler) GetUserProfileWithOrders(c *gin.Context) {
	userID := c.Param("id")
	authToken := c.GetHeader("Authorization")

	// Use aggregator service to fetch and combine data
	userProfileWithOrders, err := h.userAggregator.GetUserProfileWithOrders(c.Request.Context(), userID, authToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": userProfileWithOrders,
		"message": "User profile with orders fetched successfully",
	})
}

// GetUsersWithOrderCount fetches users with their order counts
func (h *UserAggregatorHandler) GetUsersWithOrderCount(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Use aggregator service to fetch and combine data
	usersWithOrderCount, err := h.userAggregator.GetUsersWithOrderCount(c.Request.Context(), authToken, page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": usersWithOrderCount,
		"message": "Users with order count fetched successfully",
	})
}