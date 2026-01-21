package aggregator

import (
	"context"

	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/client"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model/dto"
	"github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/proxy"
)

// CartAggregator handles cart aggregation with product data
type CartAggregator struct {
	cartClient    *client.CartClient
	productClient *client.ProductClient
	aggProxy      *proxy.AggregateProxy
}

// NewCartAggregator creates a new cart aggregator
func NewCartAggregator(
	cartClient *client.CartClient,
	productClient *client.ProductClient,
	aggProxy *proxy.AggregateProxy,
) *CartAggregator {
	return &CartAggregator{
		cartClient:    cartClient,
		productClient: productClient,
		aggProxy:      aggProxy,
	}
}

// GetCartWithProducts fetches cart with related product data
func (a *CartAggregator) GetCartWithProducts(ctx context.Context, authToken string) (*dto.CartWithProducts, error) {
	// Step 1: Cart items lao
	cartItems, err := a.cartClient.GetCartItems(ctx, authToken)
	if err != nil {
		return nil, err
	}

	// Step 2: Product IDs nikalo
	var productIDs []string
	for _, item := range cartItems {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if productID, exists := itemMap["product_id"]; exists {
				if productIDStr, ok := productID.(string); ok && productIDStr != "" {
					productIDs = append(productIDs, productIDStr)
				}
			}
		}
	}

	// Step 3: Batch mein products lao (aggregate_proxy use karke)
	products, err := a.aggProxy.BatchFetchProducts(ctx, authToken, productIDs)
	if err != nil {
		return nil, err
	}

	// Step 4: Join karo
	result := a.joinCartWithProducts(cartItems, products)

	return result, nil
}

// joinCartWithProducts joins cart items with product data
func (a *CartAggregator) joinCartWithProducts(cartItems []interface{}, products []interface{}) *dto.CartWithProducts {
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

	var enhancedCartItems []interface{}
	for _, item := range cartItems {
		if itemMap, ok := item.(map[string]interface{}); ok {
			if productID, exists := itemMap["product_id"]; exists {
				if productIDStr, ok := productID.(string); ok {
					if product, exists := productMap[productIDStr]; exists {
						// Add product details to the cart item
						itemMap["product_details"] = product
					}
				}
			}
			enhancedCartItems = append(enhancedCartItems, itemMap)
		}
	}

	return &dto.CartWithProducts{
		Cart:      enhancedCartItems,
		Products:  products,
	}
}