package aggregator

import (
	"context"
	"sync"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model/dto"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
)

// UserAggregator handles user data aggregation with related information
type UserAggregator struct {
	userClient    *client.UserClient
	orderClient   *client.OrderClient
	productClient *client.ProductClient
	cartClient    *client.CartClient
	aggProxy      *proxy.AggregateProxy
}

// NewUserAggregator creates a new user aggregator
func NewUserAggregator(
	userClient *client.UserClient,
	orderClient *client.OrderClient,
	productClient *client.ProductClient,
	cartClient *client.CartClient,
	aggProxy *proxy.AggregateProxy,
) *UserAggregator {
	return &UserAggregator{
		userClient:    userClient,
		orderClient:   orderClient,
		productClient: productClient,
		cartClient:    cartClient,
		aggProxy:      aggProxy,
	}
}

// GetUserProfileWithOrders fetches user profile with related order data
func (a *UserAggregator) GetUserProfileWithOrders(ctx context.Context, userID, authToken string) (*dto.UserProfileWithOrders, error) {
	// Step 1: User profile lao
	userProfile, err := a.userClient.GetUserByID(ctx, userID, authToken)
	if err != nil {
		return nil, err
	}

	// Step 2: User ke orders lao
	orders, err := a.orderClient.GetOrdersByUserID(ctx, userID, authToken, "1", "10")
	if err != nil {
		return nil, err
	}

	// Step 3: Extract product IDs from orders
	var productIDs []string
	for _, order := range orders {
		if orderMap, ok := order.(map[string]interface{}); ok {
			if items, exists := orderMap["items"]; exists {
				if itemsArr, ok := items.([]interface{}); ok {
					for _, item := range itemsArr {
						if itemMap, ok := item.(map[string]interface{}); ok {
							if productID, exists := itemMap["product_id"]; exists {
								if productIDStr, ok := productID.(string); ok && productIDStr != "" {
									productIDs = append(productIDs, productIDStr)
								}
							}
						}
					}
				}
			}
		}
	}

	// Step 4: Batch mein products lao (aggregate_proxy use karke)
	products, err := a.aggProxy.BatchFetchProducts(ctx, authToken, productIDs)
	if err != nil {
		return nil, err
	}

	// Step 5: Join karo
	result := &dto.UserProfileWithOrders{
		User:    userProfile,
		Orders:  orders,
		Products: products,
	}

	return result, nil
}

// GetUsersWithOrderCount fetches users with their order counts
func (a *UserAggregator) GetUsersWithOrderCount(ctx context.Context, authToken, page, limit string) ([]dto.UserWithOrderCount, error) {
	// Step 1: Users lao
	users, err := a.userClient.GetUsers(ctx, authToken, page, limit)
	if err != nil {
		return nil, err
	}

	// Step 2: User IDs nikalo
	var userIDs []string
	for _, user := range users {
		if userMap, ok := user.(map[string]interface{}); ok {
			if id, exists := userMap["id"]; exists {
				if idStr, ok := id.(string); ok && idStr != "" {
					userIDs = append(userIDs, idStr)
				}
			}
		}
	}

	// Step 3: Fetch order counts for each user
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]dto.UserWithOrderCount, len(users))
	errs := make([]error, len(users))

	for i, user := range users {
		wg.Add(1)
		go func(index int, u interface{}) {
			defer wg.Done()
			
			if userMap, ok := u.(map[string]interface{}); ok {
				if id, exists := userMap["id"]; exists {
					if userID, ok := id.(string); ok && userID != "" {
						// Fetch order count for this user
						orders, err := a.orderClient.GetOrdersByUserID(ctx, userID, authToken, "1", "1") // Just count
						if err != nil {
							mu.Lock()
							errs[index] = err
							mu.Unlock()
							return
						}
						
						var userObj dto.UserSummary
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
						if firstName, exists := userMap["first_name"]; exists {
							if firstNameStr, ok := firstName.(string); ok {
								userObj.FirstName = firstNameStr
							}
						}
						if lastName, exists := userMap["last_name"]; exists {
							if lastNameStr, ok := lastName.(string); ok {
								userObj.LastName = lastNameStr
							}
						}
						
						mu.Lock()
						results[index] = dto.UserWithOrderCount{
							User:      userObj,
							OrderCount: len(orders),
						}
						mu.Unlock()
					}
				}
			}
		}(i, user)
	}

	wg.Wait()

	// Check for any errors
	for _, err := range errs {
		if err != nil {
			return nil, err
		}
	}

	return results, nil
}