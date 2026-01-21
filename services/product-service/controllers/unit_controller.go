package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DurgaPratapRajbhar/e-commerce/product-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/product-service/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductUnitController struct {
	ProductUnitServices services.ProductUnitServices
	validator           *validator.Validate
}

func NewProductUnitController(ProductUnitServices services.ProductUnitServices) *ProductUnitController {
	return &ProductUnitController{
		ProductUnitServices: ProductUnitServices,
		validator:           validator.New(),
	}
}

// CreateUnit creates a new product Unit
// @Summary Create a new unit of measurement
// @Description Create a new unit of measurement
// @Tags Units
// @Accept json
// @Produce json
// @Param unit body models.UnitOfMeasurement true "Unit data"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-units [post]
func (pvc *ProductUnitController) CreateUnit(c *gin.Context) {
	var Unit models.UnitOfMeasurement
	if err := c.ShouldBindJSON(&Unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate input
	if err := pvc.validator.Struct(&Unit); err != nil {
		validationErrors := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, strings.Join([]string{e.Field(), e.Tag(), e.Param()}, " "))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErrors})
		return
	}

	// Call service to create product Unit
	if err := pvc.ProductUnitServices.CreateUnit(&Unit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create Unit"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product Unit created successfully", "data": Unit})
}

// GetUnit retrieves a product Unit by ID
// @Summary Get a unit by ID
// @Description Get a unit by its ID
// @Tags Units
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} models.UnitOfMeasurement
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /product-units/{id} [get]
func (pvc *ProductUnitController) GetUnit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Unit ID"})
		return
	}

	Unit, err := pvc.ProductUnitServices.GetUnit(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Unit not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Unit})
}

// GetAllUnits retrieves all Units of a product
// @Summary Get all units
// @Description Get all units of measurement
// @Tags Units
// @Accept json
// @Produce json
// @Success 200 {array} models.UnitOfMeasurement
// @Failure 500 {object} map[string]string
// @Router /product-units [get]
func (pvc *ProductUnitController) GetAllUnits(c *gin.Context) {

	Units, err := pvc.ProductUnitServices.GetAllUnit()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Units"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Units})
}

// UpdateUnit updates an existing product Unit
// @Summary Update a unit
// @Description Update a unit by its ID
// @Tags Units
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Param unit body models.UnitOfMeasurement true "Unit data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-units/{id} [put]
func (pvc *ProductUnitController) UpdateUnit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Unit ID"})
		return
	}

	var Unit models.UnitOfMeasurement
	if err := c.ShouldBindJSON(&Unit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate input
	if err := pvc.validator.Struct(&Unit); err != nil {
		validationErrors := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, strings.Join([]string{e.Field(), e.Tag(), e.Param()}, " "))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Validation failed", "details": validationErrors})
		return
	}

	// Call service to update product Unit
	if err := pvc.ProductUnitServices.UpdateUnit(uint(id), &Unit); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Unit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Unit updated successfully"})
}

// DeleteUnit deletes a product Unit by ID
// @Summary Delete a unit
// @Description Delete a unit by its ID
// @Tags Units
// @Accept json
// @Produce json
// @Param id path int true "Unit ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /product-units/{id} [delete]
func (pvc *ProductUnitController) DeleteUnit(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Unit ID"})
		return
	}

	if err := pvc.ProductUnitServices.DeleteUnit(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Unit"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product Unit deleted successfully"})
}
