package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
)

type UserProfileRepository interface {
	Create(profile *models.UserProfile) error
	GetByUserID(userID uint) (*models.UserProfile, error)
	GetByID(id uint) (*models.UserProfile, error)
	Update(userID uint, profile *models.UserProfile) error
	Delete(userID uint) error
	GetByUserIDs(userIDs []uint) ([]*models.UserProfile, error)
}