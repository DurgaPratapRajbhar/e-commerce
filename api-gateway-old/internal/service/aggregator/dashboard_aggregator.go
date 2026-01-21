package aggregator

import (
	"context"
	"sync"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model/dto"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
)

// DashboardAggregator handles dashboard data aggregation from multiple services
type DashboardAggregator struct {
	orderClient   *client.OrderClient
	userClient    *client.UserClient
	productClient *client.ProductClient
	cartClient    *client.CartClient
	aggProxy      *proxy.AggregateProxy
}

// NewDashboardAggregator creates a new dashboard aggregator
func NewDashboardAggregator(
	orderClient *client.OrderClient,
	userClient *client.UserClient,
	productClient *client.ProductClient,
	cartClient *client.CartClient,
	aggProxy *proxy.AggregateProxy,
) *DashboardAggregator {
	return &DashboardAggregator{
		orderClient:   orderClient,
		userClient:    userClient,
		productClient: productClient,
		cartClient:    cartClient,
		aggProxy:      aggProxy,
	}
}

// GetDashboardData fetches comprehensive dashboard data with related information
func (a *DashboardAggregator) GetDashboardData(ctx context.Context, authToken string) (*dto.DashboardData, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var result dto.DashboardData
	var errs []error

	// Fetch orders with users concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		ordersWithUsers, err := a.getOrdersWithUsers(ctx, authToken)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		result.OrdersWithUsers = ordersWithUsers
		mu.Unlock()
	}()

	// Fetch carts with products concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		cartsWithProducts, err := a.getCartsWithProducts(ctx, authToken)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		result.CartsWithProducts = cartsWithProducts
		mu.Unlock()
	}()

	// Fetch recent users concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		recentUsers, err := a.getRecentUsers(ctx, authToken)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		result.RecentUsers = recentUsers
		mu.Unlock()
	}()

	// Fetch popular products concurrently
	wg.Add(1)
	go func() {
		defer wg.Done()
		popularProducts, err := a.getPopularProducts(ctx, authToken)
		if err != nil {
			mu.Lock()
			errs = append(errs, err)
			mu.Unlock()
			return
		}
		mu.Lock()
		result.PopularProducts = popularProducts
		mu.Unlock()
	}()

	wg.Wait()

	if len(errs) > 0 {
		// Return the first error if any occurred
		return nil, errs[0]
	}

	return &result, nil
}

// getOrdersWithUsers fetches orders with related user data
func (a *DashboardAggregator) getOrdersWithUsers(ctx context.Context, authToken string) ([]dto.OrderWithUser, error) {
	// Step 1: Orders lao
	orders, err := a.orderClient.GetOrders(ctx, authToken, "1", "10") // Get recent 10 orders
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

// getCartsWithProducts fetches carts with related product data
func (a *DashboardAggregator) getCartsWithProducts(ctx context.Context, authToken string) ([]dto.CartWithProducts, error) {
	// For dashboard, we might fetch multiple user carts or sample data
	// In a real scenario, this would be more complex
	return nil, nil // Placeholder for now
}

// getRecentUsers fetches recent users
func (a *DashboardAggregator) getRecentUsers(ctx context.Context, authToken string) ([]interface{}, error) {
	// Fetch recent users from user service
	users, err := a.userClient.GetUsers(ctx, authToken, "1", "5") // Get recent 5 users
	if err != nil {
		return nil, err
	}

	return users, nil
}

// getPopularProducts fetches popular products
func (a *DashboardAggregator) getPopularProducts(ctx context.Context, authToken string) ([]interface{}, error) {
	// Fetch popular products from product service
	products, err := a.productClient.GetProducts(ctx, authToken, "1", "5") // Get top 5 products
	if err != nil {
		return nil, err
	}

	return products, nil
}

// joinOrdersWithUsers joins orders with user data
func (a *DashboardAggregator) joinOrdersWithUsers(orders []interface{}, users []interface{}) []dto.OrderWithUser {
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
			var orderObj dto.OrderSummary
			// Convert order map to OrderSummary struct
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
			if status, exists := orderMap["status"]; exists {
				if statusStr, ok := status.(string); ok {
					orderObj.Status = statusStr
				}
			}
			if totalAmount, exists := orderMap["total_amount"]; exists {
				if amount, ok := totalAmount.(float64); ok {
					orderObj.TotalAmount = amount
				} else if _, ok := totalAmount.(string); ok {
					// Handle string amounts if needed
					orderObj.TotalAmount = 0 // placeholder
				}
			}

			// Find corresponding user
			var userObj dto.UserSummary
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
				}
			}

			result = append(result, dto.OrderWithUser{
				Order: model.Order{ID: orderObj.ID, UserID: orderObj.UserID, Status: orderObj.Status, TotalPrice: orderObj.TotalAmount},
				User:  model.User{ID: userObj.ID, Email: userObj.Email, FirstName: userObj.FirstName, LastName: userObj.LastName},
			})
		}
	}

	return result
}