package dto

import (
	"github.com/DurgaPratapRajbhar/e-commerce/api-gateway/internal/model"
)

type AdminUserDTO struct {
	Auth      model.Auth          `json:"auth,omitempty"`
	Profile   model.UserProfile   `json:"profile,omitempty"`
	Addresses []model.UserAddress `json:"addresses,omitempty"`
}
