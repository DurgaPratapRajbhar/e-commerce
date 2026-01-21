package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/models"
)

type CartRepository interface {
	Create(cart *models.Cart) error
	GetByID(id uint64) (*models.Cart, error)
	GetByUserID(userID uint64) ([]models.Cart, error)
	Update(cart *models.Cart) error
	Delete(id uint64) error
	DeleteByUserID(userID uint64) error
}