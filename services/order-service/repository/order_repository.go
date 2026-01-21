package repository

import "github.com/DurgaPratapRajbhar/e-commerce/services/order-service/models"

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(orderID uint) (*models.Order, error)
	GetOrdersByUserID(userID uint) ([]models.Order, error)
	UpdateOrder(orderID uint, order *models.Order) error
	DeleteOrder(orderID uint) error
	GetAllOrders() ([]models.Order, error)
	UpdateOrderStatus(orderID uint, status string) error
	UpdatePaymentStatus(orderID uint, status string, paymentID *string) error
}
