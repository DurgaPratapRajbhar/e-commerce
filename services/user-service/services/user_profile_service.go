package services

import (
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
)

type UserProfileService interface {
	CreateProfile(profile *models.UserProfile) error
	GetProfileByUserID(userID uint) (*models.UserProfile, error)
	GetProfileByID(id uint) (*models.UserProfile, error)
	UpdateProfile(userID uint, profile *models.UserProfile) error
	DeleteProfile(userID uint) error
	GetProfilesByUserIDs(userIDs []uint) ([]*models.UserProfile, error)
}