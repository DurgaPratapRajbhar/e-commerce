package services

import (
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
)

type UserAddressService interface {
	CreateAddress(address *models.UserAddress) error
	GetAddressByID(id uint) (*models.UserAddress, error)
	GetAddressesByUserID(userID uint) ([]models.UserAddress, error)
	GetAddressByUserIDAndType(userID uint, addressType string) (*models.UserAddress, error)
	UpdateAddress(id uint, address *models.UserAddress) error
	DeleteAddress(id uint) error
	SetDefaultAddress(userID uint, addressID uint) error
	GetDefaultAddress(userID uint) (*models.UserAddress, error)
}