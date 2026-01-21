package dto

import "github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model"

// OrderDetail represents complete order details with user and product info
type OrderDetail struct {
	Order    model.Order   `json:"order"`
	User     model.User    `json:"user"`
	Product  model.Product `json:"product"`
}