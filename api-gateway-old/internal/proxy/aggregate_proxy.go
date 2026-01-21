package proxy

import (
	"context"
	"sync"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
	 
)

// AggregateProxy handles aggregate endpoints that combine data from multiple services
// Provides batch and parallel helpers for aggregators
type AggregateProxy struct {
	authClient *client.AuthClient
	userClient *client.UserClient
	orderClient *client.OrderClient
	productClient *client.ProductClient
}

// NewAggregateProxy creates a new aggregate proxy
func NewAggregateProxy(sp *ServiceProxy) *AggregateProxy {
	authClient := client.NewAuthClient(sp.config)
	userClient := client.NewUserClient(sp.config)
	orderClient := client.NewOrderClient(sp.config)
	productClient := client.NewProductClient(sp.config)
	
	return &AggregateProxy{
		authClient:    authClient,
		userClient:    userClient,
		orderClient:   orderClient,
		productClient: productClient,
	}
}

// BatchFetchUsers fetches users in batch
func (ap *AggregateProxy) BatchFetchUsers(ctx context.Context, authToken string, userIDs []string) ([]interface{}, error) {
	// Create batch request to User Service
	users, err := ap.userClient.GetProfilesByUserIDs(ctx, authToken, userIDs)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// BatchFetchProducts fetches products in batch
func (ap *AggregateProxy) BatchFetchProducts(ctx context.Context, authToken string, productIDs []string) ([]interface{}, error) {
	// Create batch request to Product Service
	products, err := ap.productClient.GetProductsByIDs(ctx, authToken, productIDs)
	if err != nil {
		return nil, err
	}
	return products, nil
}

// ParallelFetch executes multiple requests in parallel
func (ap *AggregateProxy) ParallelFetch(requests []func() (interface{}, error)) ([]interface{}, []error) {
	var wg sync.WaitGroup
	results := make([]interface{}, len(requests))
	errors := make([]error, len(requests))

	for i, req := range requests {
		wg.Add(1)
		go func(index int, request func() (interface{}, error)) {
			defer wg.Done()
			result, err := request()
			results[index] = result
			errors[index] = err
		}(i, req)
	}

	wg.Wait()
	return results, errors
}

// AggregateAdminUsers handles the admin users aggregation endpoint
func (ap *AggregateProxy) AggregateAdminUsers(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Extract token from context (set by AuthMiddleware)
	token := c.GetString("token")

	
	if token == "" {
		c.JSON(401, utils.ErrorResponse(utils.ErrUnauthorized, "Unauthorized - Missing user information", nil, utils.GenerateRequestID()))
		return
	}
	
	// Format the token for the client request
	authToken := "Bearer " + token



	// Use parallel fetch to get users from multiple services
	requests := []func() (interface{}, error){
		func() (interface{}, error) {
			return ap.authClient.GetUsers(c.Request.Context(), authToken, page, limit)
		},
	}

	results, errs := ap.ParallelFetch(requests)
	if len(errs) > 0 && errs[0] != nil {
		c.JSON(500, utils.ErrorResponse(utils.ErrInternalServer, errs[0].Error(), nil, utils.GenerateRequestID()))
		return
	}

	c.JSON(200, utils.SuccessResponse(results[0], "Users aggregated successfully", utils.GenerateRequestID()))
}