package aggregator

import (
	"net/http"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/service/aggregator"
	"github.com/gin-gonic/gin"
)

type DashboardAggregatorHandler struct {
	dashboardAggregator *aggregator.DashboardAggregator
}

func NewDashboardAggregatorHandler(
	orderClient *client.OrderClient,
	userClient *client.UserClient,
	productClient *client.ProductClient,
	cartClient *client.CartClient,
	aggProxy *proxy.AggregateProxy,
) *DashboardAggregatorHandler {
	dashboardAggregator := aggregator.NewDashboardAggregator(orderClient, userClient, productClient, cartClient, aggProxy)
	return &DashboardAggregatorHandler{
		dashboardAggregator: dashboardAggregator,
	}
}

// JOIN Operations: handler/aggregator → service/aggregator → multiple clients → combine
// Complex, slower (but optimized with batching) - gets comprehensive dashboard data
func (h *DashboardAggregatorHandler) GetDashboardData(c *gin.Context) {
	authToken := c.GetHeader("Authorization")

	// Use aggregator service to fetch and combine data from multiple services
	dashboardData, err := h.dashboardAggregator.GetDashboardData(c.Request.Context(), authToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": dashboardData,
		"message": "Dashboard data fetched successfully",
	})
}