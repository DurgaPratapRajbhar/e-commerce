package impl

import (
	"errors"
	"strings"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/services"
)

type UserProfileServiceImpl struct {
	repo repository.UserProfileRepository
}

func NewUserProfileService(repo repository.UserProfileRepository) services.UserProfileService {
	return &UserProfileServiceImpl{repo: repo}
}

func (s *UserProfileServiceImpl) CreateProfile(profile *models.UserProfile) error {
	err := s.repo.Create(profile)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") && strings.Contains(err.Error(), "uni_user_profiles_user_id") {
			return errors.New(utils.ErrAlreadyExists + ": profile already exists for this user")
		}
		return err
	}
	return nil
}

func (s *UserProfileServiceImpl) GetProfileByUserID(userID uint) (*models.UserProfile, error) {
	return s.repo.GetByUserID(userID)
}

func (s *UserProfileServiceImpl) GetProfileByID(id uint) (*models.UserProfile, error) {
	return s.repo.GetByID(id)
}

func (s *UserProfileServiceImpl) UpdateProfile(userID uint, profile *models.UserProfile) error {
	return s.repo.Update(userID, profile)
}

func (s *UserProfileServiceImpl) DeleteProfile(userID uint) error {
	return s.repo.Delete(userID)
}

func (s *UserProfileServiceImpl) GetProfilesByUserIDs(userIDs []uint) ([]*models.UserProfile, error) {
	return s.repo.GetByUserIDs(userIDs)
}

