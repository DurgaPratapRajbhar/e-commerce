package utils

import (
	"fmt"
)

// Joiner provides utilities for joining data from multiple services
type Joiner struct{}

// NewJoiner creates a new joiner utility
func NewJoiner() *Joiner {
	return &Joiner{}
}

// JoinOrdersWithUsers joins order data with user data
func (j *Joiner) JoinOrdersWithUsers(orders []interface{}, users []interface{}) []interface{} {
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

	var result []interface{}
	for _, order := range orders {
		if orderMap, ok := order.(map[string]interface{}); ok {
			if userID, exists := orderMap["user_id"]; exists {
				if userIDStr, ok := userID.(string); ok {
					if user, exists := userMap[userIDStr]; exists {
						// Add user details to the order
						orderMap["user"] = user
					}
				}
			}
			result = append(result, orderMap)
		}
	}

	return result
}

// JoinCartWithProducts joins cart items with product data
func (j *Joiner) JoinCartWithProducts(cartItems []interface{}, products []interface{}) []interface{} {
	// Create a map of products by ID for quick lookup
	productMap := make(map[string]interface{})
	for _, product := range products {
		if productMap, ok := product.(map[string]interface{}); ok {
			if id, exists := productMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					productMap[idStr] = product
				}
			}
		}
	}

	var result []interface{}
	for _, item := range cartItems {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if productID, exists := itemMap["product_id"]; exists {
				if productIDStr, ok := productID.(string); ok {
					if product, exists := productMap[productIDStr]; exists {
						// Add product details to the cart item
						itemMap["product"] = product
					}
				}
			}
			result = append(result, itemMap)
		}
	}

	return result
}

// JoinOrdersWithProducts joins order data with product data
func (j *Joiner) JoinOrdersWithProducts(orders []interface{}, products []interface{}) []interface{} {
	// Create a map of products by ID for quick lookup
	productMap := make(map[string]interface{})
	for _, product := range products {
		if productMap, ok := product.(map[string]interface{}); ok {
			if id, exists := productMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					productMap[idStr] = product
				}
			}
		}
	}

	var result []interface{}
	for _, order := range orders {
		if orderMap, ok := order.(map[string]interface{}); ok {
			// Process order items if they exist
			if items, exists := orderMap["items"]; exists {
				if itemsArr, ok := items.([]interface{}); ok {
					for _, item := range itemsArr {
						if itemMap, ok := item.(map[string]interface{}); ok {
							if productID, exists := itemMap["product_id"]; exists {
								if productIDStr, ok := productID.(string); ok {
									if product, exists := productMap[productIDStr]; exists {
										// Add product details to the order item
										itemMap["product"] = product
									}
								}
							}
						}
					}
				}
			}
			result = append(result, orderMap)
		}
	}

	return result
}

// JoinUsersWithOrders joins user data with order data
func (j *Joiner) JoinUsersWithOrders(users []interface{}, orders []interface{}) []interface{} {
	// Create a map of orders by user ID for quick lookup
	orderMap := make(map[string][]interface{})
	for _, order := range orders {
		if orderItem, ok := order.(map[string]interface{}); ok {
			if userID, exists := orderItem["user_id"]; exists {
				if userIDStr, ok := userID.(string); ok {
					// Initialize slice if it doesn't exist
					if _, exists := orderMap[userIDStr]; !exists {
						orderMap[userIDStr] = []interface{}{}
					}
					orderMap[userIDStr] = append(orderMap[userIDStr], order)
				}
			}
		}
	}

	var result []interface{}
	for _, user := range users {
		if userMap, ok := user.(map[string]interface{}); ok {
			if id, exists := userMap["id"]; exists {
				if idStr, ok := id.(string); ok {
					if userOrders, exists := orderMap[idStr]; exists {
						// Add orders to the user
						userMap["orders"] = userOrders
					}
				}
			}
			result = append(result, userMap)
		}
	}

	return result
}

// JoinMapsWithKey joins two slices of maps based on a key
func (j *Joiner) JoinMapsWithKey(primary []interface{}, secondary []interface{}, primaryJoinKey, secondaryJoinKey string) []interface{} {
	// Create a map of secondary data by join key for quick lookup
	secondaryMap := make(map[string]interface{})
	for _, item := range secondary {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if joinValue, exists := itemMap[secondaryJoinKey]; exists {
				if joinValueStr, ok := joinValue.(string); ok {
					secondaryMap[joinValueStr] = item
				}
			}
		}
	}

	var result []interface{}
	for _, primaryItem := range primary {
		if primaryMap, ok := primaryItem.(map[string]interface{}); ok {
			if joinValue, exists := primaryMap[primaryJoinKey]; exists {
				if joinValueStr, ok := joinValue.(string); ok {
					if secondaryItem, exists := secondaryMap[joinValueStr]; exists {
						// Add secondary item details to the primary item
						primaryMap[fmt.Sprintf("%s_data", secondaryJoinKey)] = secondaryItem
					}
				}
			}
			result = append(result, primaryMap)
		}
	}

	return result
}