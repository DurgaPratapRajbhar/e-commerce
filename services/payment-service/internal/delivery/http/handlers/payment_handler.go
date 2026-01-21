// Package handlers provides HTTP handlers for payment operations
package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/payment-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CreatePayment creates a new payment
// @Summary Create payment
// @Description Create a new payment
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body entity.PaymentCreateRequest true "Payment Create Request"
// @Success 200 {object} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments [post]
func CreatePayment(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.PaymentCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		payment, err := useCase.CreatePayment(c.Request.Context(), &req)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create payment", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payment, "Payment created successfully", requestID))
	}
}

// GetPayment gets a payment by ID
// @Summary Get payment
// @Description Get a payment by ID
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/{id} [get]
func GetPayment(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid payment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		payment, err := useCase.GetPayment(c.Request.Context(), uint(id))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get payment", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payment, "Payment retrieved successfully", requestID))
	}
}

// GetPaymentByOrderID gets a payment by order ID
// @Summary Get payment by order ID
// @Description Get a payment by order ID
// @Tags payments
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/order/{order_id} [get]
func GetPaymentByOrderID(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
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

		payment, err := useCase.GetPaymentByOrderID(c.Request.Context(), uint(orderID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get payment", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payment, "Payment retrieved successfully", requestID))
	}
}

// GetPaymentByTransactionID gets a payment by transaction ID
// @Summary Get payment by transaction ID
// @Description Get a payment by transaction ID
// @Tags payments
// @Accept json
// @Produce json
// @Param transaction_id path string true "Transaction ID"
// @Success 200 {object} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/transaction/{transaction_id} [get]
func GetPaymentByTransactionID(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID := c.Param("transaction_id")

		payment, err := useCase.GetPaymentByTransactionID(c.Request.Context(), transactionID)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get payment", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payment, "Payment retrieved successfully", requestID))
	}
}

// GetPaymentsByUserID gets payments by user ID
// @Summary Get payments by user ID
// @Description Get payments by user ID
// @Tags payments
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {array} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/user/{user_id} [get]
func GetPaymentsByUserID(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Param("user_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "user_id", Message: "invalid user ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		payments, err := useCase.GetPaymentsByUserID(c.Request.Context(), uint(userID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get payments", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payments, "Payments retrieved successfully", requestID))
	}
}

// UpdatePayment updates a payment
// @Summary Update payment
// @Description Update a payment
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param payment body entity.PaymentUpdateRequest true "Payment Update Request"
// @Success 200 {object} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/{id} [put]
func UpdatePayment(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid payment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		var req entity.PaymentUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		payment, err := useCase.UpdatePayment(c.Request.Context(), uint(id), &req)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update payment", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payment, "Payment updated successfully", requestID))
	}
}

// UpdatePaymentStatus updates a payment status
// @Summary Update payment status
// @Description Update a payment status
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Param status body map[string]string true "Status"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/{id}/status [patch]
func UpdatePaymentStatus(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid payment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		var req struct {
			Status string `json:"status" validate:"required,oneof=pending success failed refunded"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		err = useCase.UpdatePaymentStatus(c.Request.Context(), uint(id), req.Status)
		if err != nil {
			requestID := utils.GenerateRequestID()
			// Check if this is a validation error
			if err.Error() == "invalid payment status" {
				validationErrors := []utils.ValidationError{
					{Field: "status", Message: "invalid payment status"},
				}
				c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			} else {
				response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update payment status", err.Error(), requestID)
				c.JSON(http.StatusInternalServerError, response)
			}
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Payment status updated successfully", requestID))
	}
}

// DeletePayment deletes a payment
// @Summary Delete payment
// @Description Delete a payment
// @Tags payments
// @Accept json
// @Produce json
// @Param id path string true "Payment ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/{id} [delete]
func DeletePayment(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid payment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		err = useCase.DeletePayment(c.Request.Context(), uint(id))
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete payment", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Payment deleted successfully", requestID))
	}
}

// GetPayments gets payments with pagination
// @Summary Get payments
// @Description Get payments with pagination
// @Tags payments
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default=10
// @Param offset query int false "Offset" default=0
// @Success 200 {array} entity.Payment
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments [get]
func GetPayments(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
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

		payments, err := useCase.GetPayments(c.Request.Context(), limit, offset)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get payments", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payments, "Payments retrieved successfully", requestID))
	}
}

// GetPaymentsByStatus gets payments by status
// @Summary Get payments by status
// @Description Get payments by status
// @Tags payments
// @Accept json
// @Produce json
// @Param status path string true "Status"
// @Success 200 {array} entity.Payment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/payments/status/{status} [get]
func GetPaymentsByStatus(useCase *usecase.PaymentUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")

		payments, err := useCase.GetPaymentsByStatus(c.Request.Context(), status)
		if err != nil {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get payments", err.Error(), requestID)
			c.JSON(http.StatusInternalServerError, response)
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(payments, "Payments retrieved successfully", requestID))
	}
}