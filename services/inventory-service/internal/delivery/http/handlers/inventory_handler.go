// Package handlers provides HTTP handlers for inventory operations
package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreateInventory creates a new inventory record
// @Summary Create inventory
// @Description Create a new inventory record
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventory body entity.Inventory true "Inventory"
// @Success 200 {object} entity.Inventory
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory [post]
func CreateInventory(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inventory entity.Inventory
		if err := c.ShouldBindJSON(&inventory); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		if err := useCase.CreateInventory(c.Request.Context(), &inventory); err != nil {
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create inventory", err.Error(), utils.GenerateRequestID())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse(inventory, "Inventory created successfully", utils.GenerateRequestID()))
	}
}

// GetInventory gets inventory by product and variant ID
// @Summary Get inventory
// @Description Get inventory by product and variant ID
// @Tags inventory
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Param variant_id path string true "Variant ID"
// @Success 200 {object} entity.Inventory
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/product/{product_id}/variant/{variant_id} [get]
func GetInventory(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "product_id", Message: "invalid product ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		variantID, err := strconv.ParseUint(c.Param("variant_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "variant_id", Message: "invalid variant ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		inventory, err := useCase.GetInventory(c.Request.Context(), uint(productID), uint(variantID))
		if err != nil {
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get inventory", err.Error(), utils.GenerateRequestID())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse(inventory, "Inventory retrieved successfully", utils.GenerateRequestID()))
	}
}

// UpdateInventory updates inventory
// @Summary Update inventory
// @Description Update inventory
// @Tags inventory
// @Accept json
// @Produce json
// @Param inventory body entity.InventoryUpdateRequest true "Inventory Update Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/update [put]
func UpdateInventory(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.InventoryUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		if err := useCase.UpdateInventory(c.Request.Context(), &req); err != nil {
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update inventory", err.Error(), utils.GenerateRequestID())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse(map[string]string{"message": "inventory updated successfully"}, "Inventory updated successfully", utils.GenerateRequestID()))
	}
}

// GetLowStockItems gets low stock items
// @Summary Get low stock items
// @Description Get low stock items
// @Tags inventory
// @Accept json
// @Produce json
// @Param threshold query int false "Stock threshold" default=10
// @Success 200 {array} entity.Inventory
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/low-stock [get]
func GetLowStockItems(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		thresholdStr := c.Query("threshold")
		threshold := 10 // default threshold
		if thresholdStr != "" {
			if t, err := strconv.Atoi(thresholdStr); err == nil {
				threshold = t
			}
		}

		inventories, err := useCase.GetLowStockItems(c.Request.Context(), threshold)
		if err != nil {
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get low stock items", err.Error(), utils.GenerateRequestID())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse(inventories, "Low stock items retrieved successfully", utils.GenerateRequestID()))
	}
}

// DeleteInventory deletes inventory
// @Summary Delete inventory
// @Description Delete inventory by product and variant ID
// @Tags inventory
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Param variant_id path string true "Variant ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/product/{product_id}/variant/{variant_id} [delete]
func DeleteInventory(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		productID, err := strconv.ParseUint(c.Param("product_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "product_id", Message: "invalid product ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		variantID, err := strconv.ParseUint(c.Param("variant_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "variant_id", Message: "invalid variant ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		err = useCase.DeleteInventory(c.Request.Context(), uint(productID), uint(variantID))
		if err != nil {
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete inventory", err.Error(), utils.GenerateRequestID())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse(map[string]string{"message": "inventory deleted successfully"}, "Inventory deleted successfully", utils.GenerateRequestID()))
	}
}

// GetInventoryList gets all inventory items
// @Summary Get all inventory
// @Description Get all inventory items
// @Tags inventory
// @Accept json
// @Produce json
// @Success 200 {array} entity.Inventory
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory [get]
func GetInventoryList(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		inventories, err := useCase.GetInventoryList(c.Request.Context())
		if err != nil {
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get inventory list", err.Error(), utils.GenerateRequestID())
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		c.JSON(http.StatusOK, utils.SuccessResponse(inventories, "Low stock items retrieved successfully", utils.GenerateRequestID()))
	}
}