package impl

import (
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/database"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/repository"
)

type cartRepositoryImpl struct{}

func NewCartRepository() repository.CartRepository {
	return &cartRepositoryImpl{}
}

func (r *cartRepositoryImpl) Create(cart *models.Cart) error {
	result := database.DB.Create(cart)
	return result.Error
}

func (r *cartRepositoryImpl) GetByID(id uint64) (*models.Cart, error) {
	var cart models.Cart
	result := database.DB.First(&cart, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &cart, nil
}

func (r *cartRepositoryImpl) GetByUserID(userID uint64) ([]models.Cart, error) {
	var carts []models.Cart
	result := database.DB.Where("user_id = ?", userID).Find(&carts)
	return carts, result.Error
}

func (r *cartRepositoryImpl) Update(cart *models.Cart) error {
	result := database.DB.Save(cart)
	return result.Error
}

func (r *cartRepositoryImpl) Delete(id uint64) error {
	result := database.DB.Delete(&models.Cart{}, id)
	return result.Error
}

func (r *cartRepositoryImpl) DeleteByUserID(userID uint64) error {
	result := database.DB.Where("user_id = ?", userID).Delete(&models.Cart{})
	return result.Error
}