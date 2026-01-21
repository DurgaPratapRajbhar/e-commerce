package services

import (
	"errors"
	"fmt"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"
)

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) services.CategoryService {
	return &categoryService{repo: repo}
}

func (s *categoryService) CreateCategory(category *models.Category) error {

	if err := s.repo.CreateCategory(category); err != nil {
		fmt.Println("Category creation error:", err)
		return err
	}

	return nil

	// return s.repo.CreateCategory(category)
}

func (s *categoryService) GetCategory(id uint) (*models.Category, error) {
	return s.repo.GetCategory(id)
}

func (s *categoryService) UpdateCategory(id uint, category *models.Category) error {
	return s.repo.UpdateCategory(id, category)
}

func (s *categoryService) DeleteCategory(id uint) error {
	return s.repo.DeleteCategory(id)
}

func (s *categoryService) GetAllCategories() ([]models.Category, error) {
	return s.repo.GetAllCategories()
}

func (s *categoryService) GetAllCategoriesList() ([]models.Category, error) {
	return s.repo.GetAllCategoriesList()
}

func (s *categoryService) FindBySlugCategories(category *models.Category) error {
	var existing models.Category
	if err := s.repo.FindBySlug(existing.Slug, &existing); err == nil {
		return errors.New("category slug already exists")
	}

	return nil
}
