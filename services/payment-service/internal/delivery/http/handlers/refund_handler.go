// Package handlers provides HTTP handlers for refund operations
package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreateRefund creates a new refund
// @Summary Create refund
// @Description Create a new refund
// @Tags refunds
// @Accept json
// @Produce json
// @Param refund body entity.RefundCreateRequest true "Refund Create Request"
// @Success 200 {object} entity.Refund
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/refunds [post]
func CreateRefund(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.RefundCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		refund, err := useCase.CreateRefund(c.Request.Context(), &req)
		if err != nil {
			requestID := utils.GenerateRequestID()
			// Check if this is a validation error
			if err.Error() == "payment not found" || err.Error() == "payment must be successful to create a refund" {
				validationErrors := []utils.ValidationError{
					{Field: "payment", Message: err.Error()},
				}
				c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			} else {
				response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create refund", err.Error(), requestID)
				c.JSON(http.StatusInternalServerError, response)
			}
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(refund, "Refund created successfully", requestID))
	}
}

// GetRefund gets a refund by ID
// @Summary Get refund
// @Description Get a refund by ID
// @Tags refunds
// @Accept json
// @Produce json
// @Param id path string true "Refund ID"
// @Success 200 {object} entity.Refund
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/refunds/{id} [get]
func GetRefund(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid refund ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		refund, err := useCase.GetRefund(c.Request.Context(), uint(id))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get refund", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(refund, "Refund retrieved successfully", requestID))
	}
}

// GetRefundsByPaymentID gets refunds by payment ID
// @Summary Get refunds by payment ID
// @Description Get refunds by payment ID
// @Tags refunds
// @Accept json
// @Produce json
// @Param payment_id path string true "Payment ID"
// @Success 200 {array} entity.Refund
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/refunds/payment/{payment_id} [get]
func GetRefundsByPaymentID(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		paymentID, err := strconv.ParseUint(c.Param("payment_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "payment_id", Message: "invalid payment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		refunds, err := useCase.GetRefundsByPaymentID(c.Request.Context(), uint(paymentID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get refunds", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(refunds, "Refunds retrieved successfully", requestID))
	}
}

// GetRefundsByOrderID gets refunds by order ID
// @Summary Get refunds by order ID
// @Description Get refunds by order ID
// @Tags refunds
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {array} entity.Refund
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/refunds/order/{order_id} [get]
func GetRefundsByOrderID(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		orderID, err := strconv.ParseUint(c.Param("order_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "order_id", Message: "invalid order ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		refunds, err := useCase.GetRefundsByOrderID(c.Request.Context(), uint(orderID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get refunds", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(refunds, "Refunds retrieved successfully", requestID))
	}
}

// GetRefunds gets refunds with pagination
// @Summary Get refunds
// @Description Get refunds with pagination
// @Tags refunds
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default=10
// @Param offset query int false "Offset" default=0
// @Success 200 {array} entity.Refund
// @Failure 500 {object} map[string]string
// @Router /api/v1/refunds [get]
func GetRefunds(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		limitStr := c.Query("limit")
		offsetStr := c.Query("offset")
		
		limit := 10
		if limitStr != "" {
			if l, err := strconv.Atoi(limitStr); err == nil {
				limit = l
			}
		}
		
		offset := 0
		if offsetStr != "" {
			if o, err := strconv.Atoi(offsetStr); err == nil {
				offset = o
			}
		}

		refunds, err := useCase.GetRefunds(c.Request.Context(), limit, offset)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get refunds", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(refunds, "Refunds retrieved successfully", requestID))
	}
}

// GetRefundsByStatus gets refunds by status
// @Summary Get refunds by status
// @Description Get refunds by status
// @Tags refunds
// @Accept json
// @Produce json
// @Param status path string true "Status"
// @Success 200 {array} entity.Refund
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/refunds/status/{status} [get]
func GetRefundsByStatus(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")

		refunds, err := useCase.GetRefundsByStatus(c.Request.Context(), status)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get refunds", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(refunds, "Refunds retrieved successfully", requestID))
	}
}