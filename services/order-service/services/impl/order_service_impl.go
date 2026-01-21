package impl

import (
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/services"
)

type OrderServiceImpl struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) services.OrderService {
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) CreateOrder(order *models.Order) error {
	return s.repo.CreateOrder(order)
}

func (s *OrderServiceImpl) GetOrderByID(orderID uint) (*models.Order, error) {
	return s.repo.GetOrderByID(orderID)
}

func (s *OrderServiceImpl) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	return s.repo.GetOrdersByUserID(userID)
}

func (s *OrderServiceImpl) UpdateOrder(orderID uint, order *models.Order) error {
	return s.repo.UpdateOrder(orderID, order)
}

func (s *OrderServiceImpl) DeleteOrder(orderID uint) error {
	return s.repo.DeleteOrder(orderID)
}

func (s *OrderServiceImpl) GetAllOrders() ([]models.Order, error) {
	return s.repo.GetAllOrders()
}

func (s *OrderServiceImpl) UpdateOrderStatus(orderID uint, status string) error {
	return s.repo.UpdateOrderStatus(orderID, status)
}

func (s *OrderServiceImpl) UpdatePaymentStatus(orderID uint, status string, paymentID *string) error {
	return s.repo.UpdatePaymentStatus(orderID, status, paymentID)
}
