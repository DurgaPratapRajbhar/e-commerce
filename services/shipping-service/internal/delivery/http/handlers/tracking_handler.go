// Package handlers provides HTTP handlers for tracking operations
package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/usecase"
	"github.com/DurgaPratapRajbhar/e-commerce/shipping-service/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

// CreateTrackingEvent creates a new tracking event
// @Summary Create tracking event
// @Description Create a new tracking event
// @Tags tracking
// @Accept json
// @Produce json
// @Param tracking_event body entity.TrackingEventCreateRequest true "Tracking Event Create Request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tracking [post]
func CreateTrackingEvent(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req entity.TrackingEventCreateRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			validationErrors := utils.ParseValidationErrors(err.Error())
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		err := useCase.CreateTrackingEvent(c.Request.Context(), &req)
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to create tracking event", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Tracking event created successfully", requestID))
	}
}

// GetTrackingEvents gets tracking events for a shipment
// @Summary Get tracking events
// @Description Get tracking events for a shipment
// @Tags tracking
// @Accept json
// @Produce json
// @Param shipment_id path string true "Shipment ID"
// @Success 200 {array} entity.TrackingEvent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tracking/{shipment_id} [get]
func GetTrackingEvents(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		shipmentID, err := strconv.ParseUint(c.Param("shipment_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "shipment_id", Message: "invalid shipment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		events, err := useCase.GetTrackingEvents(c.Request.Context(), uint(shipmentID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get tracking events", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(events, "Tracking events retrieved successfully", requestID))
	}
}

// GetLatestTrackingEvent gets the latest tracking event for a shipment
// @Summary Get latest tracking event
// @Description Get the latest tracking event for a shipment
// @Tags tracking
// @Accept json
// @Produce json
// @Param shipment_id path string true "Shipment ID"
// @Success 200 {object} entity.TrackingEvent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/v1/tracking/{shipment_id}/latest [get]
func GetLatestTrackingEvent(useCase *usecase.ShippingUseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		shipmentID, err := strconv.ParseUint(c.Param("shipment_id"), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "shipment_id", Message: "invalid shipment ID"},
			}
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}

		event, err := useCase.GetLatestTrackingEvent(c.Request.Context(), uint(shipmentID))
		if err != nil {
			requestID := utils.GenerateRequestID()
			c.JSON(http.StatusInternalServerError, utils.ErrorResponse(utils.ErrInternalServer, "Failed to get latest tracking event", err.Error(), requestID))
			return
		}

		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusOK, utils.SuccessResponse(event, "Latest tracking event retrieved successfully", requestID))
	}
}