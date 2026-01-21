package aggregator

import (
	"context"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model/dto"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
)

// OrderAggregator handles order aggregation with related data
type OrderAggregator struct {
	orderClient   *client.OrderClient
	userClient  *client.UserClient
	productClient *client.ProductClient
	aggProxy    *proxy.AggregateProxy
}

// NewOrderAggregator creates a new order aggregator
func NewOrderAggregator(
	orderClient *client.OrderClient,
	userClient *client.UserClient,
	productClient *client.ProductClient,
	aggProxy *proxy.AggregateProxy,
) *OrderAggregator {
	return &OrderAggregator{
		orderClient:   orderClient,
		userClient:   userClient,
		productClient: productClient,
		aggProxy:     aggProxy,
	}
}

// GetOrdersWithUsers fetches orders with related user data
func (a *OrderAggregator) GetOrdersWithUsers(ctx context.Context, authToken, page, limit string) ([]dto.OrderWithUser, error) {
	// Step 1: Orders lao
	orders, err := a.orderClient.GetOrders(ctx, authToken, page, limit)
	if err != nil {
		return nil, err
	}

	// Step 2: User IDs nikalo
	var userIDs []string
	for _, order := range orders {
		if orderMap, ok := order.(map[string]interface{}); ok {
			if userID, exists := orderMap["user_id"]; exists {
				if userIDStr, ok := userID.(string); ok && userIDStr != "" {
					userIDs = append(userIDs, userIDStr)
				}
			}
		}
	}

	// Step 3: Batch mein users lao (aggregate_proxy use karke)
	users, err := a.aggProxy.BatchFetchUsers(ctx, authToken, userIDs)
	if err != nil {
		return nil, err
	}

	// Step 4: Join karo
	result := a.joinOrdersWithUsers(orders, users)

	return result, nil
}

// joinOrdersWithUsers joins orders with user data
func (a *OrderAggregator) joinOrdersWithUsers(orders []interface{}, users []interface{}) []dto.OrderWithUser {
	// Create a map of users by ID for quick lookup
	userMap := make(map[string]interface{})
	for _, user := range users {
		if userMap, ok := user.(map[string]interface{}); ok {
			if id, exists := userMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					userMap[idStr] = user
				}
			}
		}
	}

	var result []dto.OrderWithUser
	for _, order := range orders {
		if orderMap, ok := order.(map[string]interface{}); ok {
			var orderObj model.Order
			// Convert order map to Order struct
			if id, exists := orderMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					orderObj.ID = idStr
				}
			}
			if userID, exists := orderMap["user_id"]; exists {
				if userIDStr, ok := userID.(string); ok {
					orderObj.UserID = userIDStr
				}
			}
			// ... populate other fields

			// Find corresponding user
			var userObj model.User
			if user, exists := userMap[orderObj.UserID]; exists {
				if userMap, ok := user.(map[string]interface{}); ok {
					if id, exists := userMap["id"]; exists {
						if idStr, ok := id.(string); ok {
							userObj.ID = idStr
						}
					}
					if email, exists := userMap["email"]; exists {
						if emailStr, ok := email.(string); ok {
							userObj.Email = emailStr
						}
					}
					// ... populate other user fields
				}
			}

			result = append(result, dto.OrderWithUser{
				Order: orderObj,
				User:  userObj,
			})
		}
	}

	return result
}