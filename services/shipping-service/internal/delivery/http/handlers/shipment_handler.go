// Package handlers provides HTTP handlers for shipping operations
package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

// CreateShipment creates a new shipment
// @Summary Create shipment
// @Description Create a new shipment
// @Tags shipments
// @Accept json
// @Produce json
// @Param shipment body entity.ShipmentCreateRequest true "Shipment Create Request"
// @Success 200 {object} entity.Shipment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments [post]
func CreateShipment(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.ShipmentCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		shipment, err := useCase.CreateShipment(c.Request.Context(), &req)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to create shipment", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipment, "Shipment created successfully", requestID))
	}
}

// GetShipment gets a shipment by ID
// @Summary Get shipment
// @Description Get a shipment by ID
// @Tags shipments
// @Accept json
// @Produce json
// @Param id path string true "Shipment ID"
// @Success 200 {object} entity.Shipment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/{id} [get]
func GetShipment(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid shipment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		shipment, err := useCase.GetShipment(c.Request.Context(), uint(id))
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get shipment", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipment, "Shipment retrieved successfully", requestID))
	}
}

// GetShipmentByOrderID gets a shipment by order ID
// @Summary Get shipment by order ID
// @Description Get a shipment by order ID
// @Tags shipments
// @Accept json
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} entity.Shipment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/order/{order_id} [get]
func GetShipmentByOrderID(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
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

		shipment, err := useCase.GetShipmentByOrderID(c.Request.Context(), uint(orderID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get shipment", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipment, "Shipment retrieved successfully", requestID))
	}
}

// GetShipmentByTrackingNumber gets a shipment by tracking number
// @Summary Get shipment by tracking number
// @Description Get a shipment by tracking number
// @Tags shipments
// @Accept json
// @Produce json
// @Param tracking_number path string true "Tracking Number"
// @Success 200 {object} entity.Shipment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/tracking/{tracking_number} [get]
func GetShipmentByTrackingNumber(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		trackingNumber := c.Param("tracking_number")

		shipment, err := useCase.GetShipmentByTrackingNumber(c.Request.Context(), trackingNumber)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get shipment", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipment, "Shipment retrieved successfully", requestID))
	}
}

// UpdateShipment updates a shipment
// @Summary Update shipment
// @Description Update a shipment
// @Tags shipments
// @Accept json
// @Produce json
// @Param id path string true "Shipment ID"
// @Param shipment body entity.ShipmentUpdateRequest true "Shipment Update Request"
// @Success 200 {object} entity.Shipment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/{id} [put]
func UpdateShipment(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid shipment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		var req entity.ShipmentUpdateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		shipment, err := useCase.UpdateShipment(c.Request.Context(), uint(id), &req)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to update shipment", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipment, "Shipment updated successfully", requestID))
	}
}

// UpdateShipmentStatus updates a shipment status
// @Summary Update shipment status
// @Description Update a shipment status
// @Tags shipments
// @Accept json
// @Produce json
// @Param id path string true "Shipment ID"
// @Param status body map[string]string true "Status"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/{id}/status [patch]
func UpdateShipmentStatus(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid shipment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		var req struct {
			Status string `json:"status" validate:"required,oneof=pending in_transit delivered returned"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		err = useCase.UpdateShipmentStatus(c.Request.Context(), uint(id), req.Status)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to update shipment status", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Shipment status updated successfully", requestID))
	}
}

// DeleteShipment deletes a shipment
// @Summary Delete shipment
// @Description Delete a shipment
// @Tags shipments
// @Accept json
// @Produce json
// @Param id path string true "Shipment ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/{id} [delete]
func DeleteShipment(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "id", Message: "invalid shipment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		err = useCase.DeleteShipment(c.Request.Context(), uint(id))
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete shipment", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Shipment deleted successfully", requestID))
	}
}

// GetShipments gets shipments with pagination
// @Summary Get shipments
// @Description Get shipments with pagination
// @Tags shipments
// @Accept json
// @Produce json
// @Param limit query int false "Limit" default=10
// @Param offset query int false "Offset" default=0
// @Success 200 {array} entity.Shipment
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments [get]
func GetShipments(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
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

		shipments, err := useCase.GetShipments(c.Request.Context(), limit, offset)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get shipments", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipments, "Shipments retrieved successfully", requestID))
	}
}

// GetShipmentsByStatus gets shipments by status
// @Summary Get shipments by status
// @Description Get shipments by status
// @Tags shipments
// @Accept json
// @Produce json
// @Param status path string true "Status"
// @Success 200 {array} entity.Shipment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/shipments/status/{status} [get]
func GetShipmentsByStatus(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		status := c.Param("status")

		shipments, err := useCase.GetShipmentsByStatus(c.Request.Context(), status)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get shipments by status", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(shipments, "Shipments retrieved successfully", requestID))
	}
}