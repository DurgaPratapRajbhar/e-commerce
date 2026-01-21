package repository

import (
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &CategoryRepository{db: db}
}

func (r *CategoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *CategoryRepository) GetCategory(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id=?", id).First(&category, id).Error
	return &category, err
}

func (r *CategoryRepository) UpdateCategory(id uint, category *models.Category) error {

	return r.db.Model(&models.Category{}).Where("id = ?", id).Updates(category).Error
}

func (r *CategoryRepository) DeleteCategory(id uint) error {
	return r.db.Delete(&models.Category{}, id).Error
}

func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *CategoryRepository) GetAllCategoriesList() ([]models.Category, error) {
	var categories []models.Category

	// Fetch only root categories (where ParentID is 0) and preload children recursively
	err := r.db.Where("parent_id = ?", 0).
		Preload("Children.Children"). // Recursively load subcategories
		Find(&categories).Error

	return categories, err
}

// func (r *CategoryRepository) GetAllCategories() ([]models.Category, error) {
// 	var categories []models.Category

// 	// Fetch only root categories where ParentID = 0 and status = 'active'
// 	err := r.db.Where("parent_id = ? AND status = ?", 0, "active").
// 		Preload("Children", "status = ?", "active").          // Preload only active children
// 		Preload("Children.Children", "status = ?", "active"). // Recursive loading of active children
// 		Find(&categories).Error

// 	return categories, err
// }

func (r *CategoryRepository) FindBySlug(slug string, category *models.Category) error {
	return r.db.Where("slug = ?", slug).First(category).Error
}
