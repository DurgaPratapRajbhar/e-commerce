package services

import (
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/models"
)

type CartService interface {
	AddToCart(cart *models.Cart) error
	GetCartByID(id uint64) (*models.Cart, error)
	GetCartByUserID(userID uint64) ([]models.Cart, error)
	UpdateCart(cart *models.Cart) error
	RemoveFromCart(id uint64) error
	ClearCart(userID uint64) error
}