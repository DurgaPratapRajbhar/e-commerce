package impl

import (
	"errors"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/repository"

	"gorm.io/gorm"
)

type OrderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &OrderRepositoryImpl{db: db}
}

func (r *OrderRepositoryImpl) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepositoryImpl) GetOrderByID(orderID uint) (*models.Order, error) {
	var order models.Order
	if err := r.db.Preload("OrderItems").First(&order, orderID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return &order, nil
}

func (r *OrderRepositoryImpl) GetOrdersByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Preload("OrderItems").Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) UpdateOrder(orderID uint, order *models.Order) error {
	result := r.db.Model(&models.Order{}).Where("id = ?", orderID).Updates(order)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func (r *OrderRepositoryImpl) DeleteOrder(orderID uint) error {
	result := r.db.Delete(&models.Order{}, orderID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func (r *OrderRepositoryImpl) GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := r.db.Preload("OrderItems").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (r *OrderRepositoryImpl) UpdateOrderStatus(orderID uint, status string) error {
	result := r.db.Model(&models.Order{}).Where("id = ?", orderID).Update("status", status)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}

func (r *OrderRepositoryImpl) UpdatePaymentStatus(orderID uint, status string, paymentID *string) error {
	updates := map[string]interface{}{
		"payment_status": status,
	}
	if paymentID != nil {
		updates["payment_id"] = paymentID
	}

	result := r.db.Model(&models.Order{}).Where("id = ?", orderID).Updates(updates)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("order not found")
	}
	return nil
}
