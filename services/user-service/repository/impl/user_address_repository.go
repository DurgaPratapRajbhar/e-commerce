package impl

import (
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/repository"

	"gorm.io/gorm"
)

type UserAddressRepositoryImpl struct {
	db *gorm.DB
}

func NewUserAddressRepository(db *gorm.DB) repository.UserAddressRepository {
	return &UserAddressRepositoryImpl{db: db}
}

func (r *UserAddressRepositoryImpl) Create(address *models.UserAddress) error {
	return r.db.Create(address).Error
}

func (r *UserAddressRepositoryImpl) GetByID(id uint) (*models.UserAddress, error) {
	var address models.UserAddress
	err := r.db.First(&address, id).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *UserAddressRepositoryImpl) GetByUserID(userID uint) ([]models.UserAddress, error) {
	var addresses []models.UserAddress
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	if err != nil {
		return nil, err
	}
	return addresses, nil
}

func (r *UserAddressRepositoryImpl) GetByUserIDAndType(userID uint, addressType string) (*models.UserAddress, error) {
	var address models.UserAddress
	err := r.db.Where("user_id = ? AND address_type = ?", userID, addressType).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}

func (r *UserAddressRepositoryImpl) Update(id uint, address *models.UserAddress) error {
	return r.db.Model(&models.UserAddress{}).Where("id = ?", id).Updates(address).Error
}

func (r *UserAddressRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&models.UserAddress{}, id).Error
}

func (r *UserAddressRepositoryImpl) SetDefaultAddress(userID uint, addressID uint) error {
	if err := r.db.Model(&models.UserAddress{}).
		Where("user_id = ? AND is_default = ?", userID, true).
		Update("is_default", false).Error; err != nil {
		return err
	}

	return r.db.Model(&models.UserAddress{}).
		Where("id = ? AND user_id = ?", addressID, userID).
		Update("is_default", true).Error
}

func (r *UserAddressRepositoryImpl) GetDefaultAddress(userID uint) (*models.UserAddress, error) {
	var address models.UserAddress
	err := r.db.Where("user_id = ? AND is_default = ?", userID, true).First(&address).Error
	if err != nil {
		return nil, err
	}
	return &address, nil
}