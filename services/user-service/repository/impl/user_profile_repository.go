package impl

import (
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/repository"

	"gorm.io/gorm"
)

type UserProfileRepositoryImpl struct {
	db *gorm.DB
}

func NewUserProfileRepository(db *gorm.DB) repository.UserProfileRepository {
	return &UserProfileRepositoryImpl{db: db}
}

func (r *UserProfileRepositoryImpl) Create(profile *models.UserProfile) error {
	return r.db.Create(profile).Error
}

func (r *UserProfileRepositoryImpl) GetByUserID(userID uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.Where("user_id = ?", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *UserProfileRepositoryImpl) GetByID(id uint) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := r.db.First(&profile, id).Error
	if err != nil {
		return nil, err
	}
	return &profile, nil
}

func (r *UserProfileRepositoryImpl) Update(userID uint, profile *models.UserProfile) error {
	return r.db.Model(&models.UserProfile{}).Where("user_id = ?", userID).Updates(profile).Error
}

func (r *UserProfileRepositoryImpl) Delete(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.UserProfile{}).Error
}

func (r *UserProfileRepositoryImpl) GetByUserIDs(userIDs []uint) ([]*models.UserProfile, error) {
	var profiles []*models.UserProfile
	if len(userIDs) == 0 {
		return profiles, nil
	}
	
	// Create a query with placeholders for the IN clause
	err := r.db.Where("user_id IN ?", userIDs).Find(&profiles).Error
	if err != nil {
		return nil, err
	}
	return profiles, nil
}

