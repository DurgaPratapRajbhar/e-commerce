// Package handlers provides HTTP handlers for inventory transaction operations
package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/inventory-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreateTransaction creates a new inventory transaction
// @Summary Create inventory transaction
// @Description Create a new inventory transaction
// @Tags inventory-transactions
// @Accept json
// @Produce json
// @Param transaction body entity.InventoryTransactionRequest true "Inventory Transaction Request"
// @Success 200 {object} entity.InventoryTransaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/transactions [post]
func CreateTransaction(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.InventoryTransactionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		// For now, we'll just create an inventory update request
		updateReq := &entity.InventoryUpdateRequest{
			ProductID:       req.ProductID,
			VariantID:       req.VariantID,
			QuantityChange:  req.Quantity,
			TransactionType: req.TransactionType,
			ReferenceID:     req.ReferenceID,
		}

		if req.TransactionType == "out" || req.TransactionType == "reserved" {
			updateReq.QuantityChange = -req.Quantity // Negative for out/reserved
		}

		if err := useCase.UpdateInventory(c.Request.Context(), updateReq); err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to process transaction", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		// Return success message
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(map[string]string{"message": "transaction processed successfully"}, "Transaction processed successfully", requestID))
	}
}

// GetTransactionsByProduct gets transactions by product ID
// @Summary Get transactions by product
// @Description Get inventory transactions by product ID
// @Tags inventory-transactions
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Success 200 {array} entity.InventoryTransaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/transactions/product/{product_id} [get]
func GetTransactionsByProduct(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
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

		transactions, err := useCase.GetInventoryTransactions(c.Request.Context(), uint(productID), 0)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get transactions", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(transactions, "Transactions retrieved successfully", requestID))
	}
}

// GetTransactionsByProductAndVariant gets transactions by product and variant ID
// @Summary Get transactions by product and variant
// @Description Get inventory transactions by product and variant ID
// @Tags inventory-transactions
// @Accept json
// @Produce json
// @Param product_id path string true "Product ID"
// @Param variant_id path string true "Variant ID"
// @Success 200 {array} entity.InventoryTransaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/transactions/product/{product_id}/variant/{variant_id} [get]
func GetTransactionsByProductAndVariant(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
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

		transactions, err := useCase.GetInventoryTransactions(c.Request.Context(), uint(productID), uint(variantID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get transactions", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(transactions, "Transactions retrieved successfully", requestID))
	}
}

// GetTransactionsByReference gets transactions by reference ID
// @Summary Get transactions by reference
// @Description Get inventory transactions by reference ID
// @Tags inventory-transactions
// @Accept json
// @Produce json
// @Param reference_id path string true "Reference ID"
// @Success 200 {array} entity.InventoryTransaction
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/transactions/reference/{reference_id} [get]
func GetTransactionsByReference(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		referenceID, err := strconv.ParseUint(c.Param("reference_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "reference_id", Message: "invalid reference ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		// Note: This functionality is not implemented in the usecase yet
		// We'll return a placeholder response
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(map[string]interface{}{"reference_id": referenceID}, "Not implemented yet", requestID))
	}
}

// GetRecentTransactions gets recent transactions
// @Summary Get recent transactions
// @Description Get recent inventory transactions
// @Tags inventory-transactions
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default=10
// @Success 200 {array} entity.InventoryTransaction
// @Failure 500 {object} map[string]string
// @Router /api/v1/inventory/transactions/recent [get]
func GetRecentTransactions(useCase *usecase.InventoryUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.Query("limit")
		limit := 10 // default limit
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}

		transactions, err := useCase.GetRecentTransactions(c.Request.Context(), limit)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get recent transactions", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(transactions, "Recent transactions retrieved successfully", requestID))
	}
}