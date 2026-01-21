package impl

import (
	"errors"
	"strings"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/services"
)

type UserAddressServiceImpl struct {
	repo repository.UserAddressRepository
}

func NewUserAddressService(repo repository.UserAddressRepository) services.UserAddressService {
	return &UserAddressServiceImpl{repo: repo}
}

func (s *UserAddressServiceImpl) CreateAddress(address *models.UserAddress) error {
	err := s.repo.Create(address)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(utils.ErrAlreadyExists + ": address already exists")
		}
		return err
	}
	return nil
}

func (s *UserAddressServiceImpl) GetAddressByID(id uint) (*models.UserAddress, error) {
	return s.repo.GetByID(id)
}

func (s *UserAddressServiceImpl) GetAddressesByUserID(userID uint) ([]models.UserAddress, error) {
	return s.repo.GetByUserID(userID)
}

func (s *UserAddressServiceImpl) GetAddressByUserIDAndType(userID uint, addressType string) (*models.UserAddress, error) {
	return s.repo.GetByUserIDAndType(userID, addressType)
}

func (s *UserAddressServiceImpl) UpdateAddress(id uint, address *models.UserAddress) error {
	return s.repo.Update(id, address)
}

func (s *UserAddressServiceImpl) DeleteAddress(id uint) error {
	return s.repo.Delete(id)
}

func (s *UserAddressServiceImpl) SetDefaultAddress(userID uint, addressID uint) error {
	return s.repo.SetDefaultAddress(userID, addressID)
}

func (s *UserAddressServiceImpl) GetDefaultAddress(userID uint) (*models.UserAddress, error) {
	return s.repo.GetDefaultAddress(userID)
}