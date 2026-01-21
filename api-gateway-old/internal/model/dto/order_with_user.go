package dto

import "github.com/DurgaPratapRajbhar/ecommerce-microservices/api-gateway/internal/model"

// OrderWithUser represents an order with user details
type OrderWithUser struct {
	Order      model.Order `json:"order"`
	User       model.User  `json:"user"`
}