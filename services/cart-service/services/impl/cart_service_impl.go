package impl

import (
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/services"
)

type cartServiceImpl struct {
	cartRepository repository.CartRepository
}

func NewCartService(repo repository.CartRepository) services.CartService {
	return &cartServiceImpl{
		cartRepository: repo,
	}
}

func (s *cartServiceImpl) AddToCart(cart *models.Cart) error {
	return s.cartRepository.Create(cart)
}

func (s *cartServiceImpl) GetCartByID(id uint64) (*models.Cart, error) {
	return s.cartRepository.GetByID(id)
}

func (s *cartServiceImpl) GetCartByUserID(userID uint64) ([]models.Cart, error) {
	return s.cartRepository.GetByUserID(userID)
}

func (s *cartServiceImpl) UpdateCart(cart *models.Cart) error {
	return s.cartRepository.Update(cart)
}

func (s *cartServiceImpl) RemoveFromCart(id uint64) error {
	return s.cartRepository.Delete(id)
}

func (s *cartServiceImpl) ClearCart(userID uint64) error {
	return s.cartRepository.DeleteByUserID(userID)
}