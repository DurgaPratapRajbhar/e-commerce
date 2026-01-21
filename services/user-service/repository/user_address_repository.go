package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
)

type UserAddressRepository interface {
	Create(address *models.UserAddress) error
	GetByID(id uint) (*models.UserAddress, error)
	GetByUserID(userID uint) ([]models.UserAddress, error)
	GetByUserIDAndType(userID uint, addressType string) (*models.UserAddress, error)
	Update(id uint, address *models.UserAddress) error
	Delete(id uint) error
	SetDefaultAddress(userID uint, addressID uint) error
	GetDefaultAddress(userID uint) (*models.UserAddress, error)
}