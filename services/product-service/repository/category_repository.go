package repository

import "github.com/DurgaPratapRajbhar/e-commerce/product-service/models"

type CategoryRepository interface {
	CreateCategory(category *models.Category) error
	GetCategory(id uint) (*models.Category, error)
	UpdateCategory(id uint, category *models.Category) error
	DeleteCategory(id uint) error
	GetAllCategories() ([]models.Category, error)
	FindBySlug(slug string, category *models.Category) error
	GetAllCategoriesList() ([]models.Category, error)
}
