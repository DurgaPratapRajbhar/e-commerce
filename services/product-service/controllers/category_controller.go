package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

// CreateCategory godoc
// @Summary Create a new category
// @Description Create a new category with optional image upload
// @Tags categories
// @Accept json
// @Produce json
// @Param category body models.Category true "Category data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /categories [post]
func (cc *CategoryController) CreateCategory(c *gin.Context) {
	// Initialize Category
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		var errorMessages []map[string]string
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errorMessages = append(errorMessages, map[string]string{
					"field":   e.Field(),
					"message": fmt.Sprintf("The %s field is %s", e.Field(), e.Tag()),
				})
			}
		}

		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input data", errorMessages, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Validation
	validate := validator.New()
	if err := validate.Struct(&category); err != nil {
		var errorMessages []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, map[string]string{
				"field":   e.Field(),
				"message": fmt.Sprintf("The %s field is %s", e.Field(), e.Tag()),
			})
		}

		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input data", errorMessages, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Check if categoryService is not nil
	if cc.categoryService == nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Category service not initialized", "Category service is not initialized", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if err := cc.categoryService.CreateCategory(&category); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create category", "Failed to create category", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Successful response
	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusCreated, utils.SuccessResponse(category, "Category created successfully", requestID))
}

// GetCategory godoc
// @Summary Get a category by ID
// @Description Get a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /categories/{id} [get]
func (cc *CategoryController) GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid category ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	category, err := cc.categoryService.GetCategory(uint(id))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrNotFound, "Category not found", "Category not found", requestID)
		c.JSON(http.StatusNotFound, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(category, "Category retrieved successfully", requestID))
}

// GetAllCategories godoc
// @Summary Get all categories
// @Description Get all categories
// @Tags categories
// @Accept json
// @Produce json
// @Success 200 {object} string
// @Failure 500 {object} map[string]string
// @Router /categories [get]
func (cc *CategoryController) GetAllCategories(c *gin.Context) {
	categories, err := cc.categoryService.GetAllCategories()
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to fetch categories", "Failed to fetch categories", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	fmt.Println(categories)

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(categories, "Categories retrieved successfully", requestID))
}

func (cc *CategoryController) GetAllCategoriesList(c *gin.Context) {
	categories, err := cc.categoryService.GetAllCategoriesList()
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to fetch categories", "Failed to fetch categories", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	fmt.Println(categories)

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(categories, "Categories retrieved successfully", requestID))
}

// UpdateCategory godoc
// @Summary Update a category
// @Description Update a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param category body models.Category true "Category data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [put]
func (cc *CategoryController) UpdateCategory(c *gin.Context) {

	cat_id := c.Param("id")
	id, err := strconv.ParseUint(cat_id, 10, 64)
	if err != nil || id == 0 {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid category ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		var errorMessages []map[string]string
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errorMessages = append(errorMessages, map[string]string{
					"field":   e.Field(),
					"message": fmt.Sprintf("The %s field is %s", e.Field(), e.Tag()),
				})
			}
		}

		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input data", errorMessages, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	validate := validator.New()
	if err := validate.Struct(&category); err != nil {
		var errorMessages []map[string]string
		for _, e := range err.(validator.ValidationErrors) {
			errorMessages = append(errorMessages, map[string]string{
				"field":   e.Field(),
				"message": fmt.Sprintf("The %s field is %s", e.Field(), e.Tag()),
			})
		}

		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input data", errorMessages, requestID)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	err = cc.categoryService.UpdateCategory(uint(id), &category)
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update category", "Failed to Update category", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusCreated, utils.SuccessResponse(category, "Category Updated successfully", requestID))

}

// DeleteCategory godoc
// @Summary Delete a category
// @Description Delete a category by its ID
// @Tags categories
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Success 200 {object} string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /categories/{id} [delete]
func (cc *CategoryController) DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid category ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := cc.categoryService.DeleteCategory(uint(id)); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete category", "Failed to delete category", requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Category deleted successfully", requestID))
}
