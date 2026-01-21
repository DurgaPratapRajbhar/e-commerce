package handlers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/services/order-service/services"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service services.OrderService
}

func NewOrderHandler(service services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.Order true "Order data"
// @Success 201 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := h.service.CreateOrder(&order); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create order", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusCreated, utils.SuccessResponse(order, "Order created successfully", requestID))
}

// GetOrder godoc
// @Summary Get an order by ID
// @Description Get an order by its ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.Order
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid order ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	order, err := h.service.GetOrderByID(uint(id))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrNotFound, "Order not found", err.Error(), requestID)
		c.JSON(http.StatusNotFound, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(order, "Order retrieved successfully", requestID))
}

// GetUserOrders godoc
// @Summary Get all orders by user ID
// @Description Get all orders by user ID
// @Tags orders
// @Accept json
// @Produce json
// @Param userId path int true "User ID"
// @Success 200 {array} models.Order
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/user/{userId} [get]
func (h *OrderHandler) GetUserOrders(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "userId", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	orders, err := h.service.GetOrdersByUserID(uint(userID))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to retrieve orders", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(orders, "Orders retrieved successfully", requestID))
}

// GetAllOrders godoc
// @Summary Get all orders
// @Description Get all orders
// @Tags orders
// @Accept json
// @Produce json
// @Success 200 {array} models.Order
// @Failure 500 {object} map[string]string
// @Router /orders [get]
func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	orders, err := h.service.GetAllOrders()
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to retrieve orders", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(orders, "Orders retrieved successfully", requestID))
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param order body models.Order true "Order data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid order ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := h.service.UpdateOrder(uint(id), &order); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update order", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(order, "Order updated successfully", requestID))
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id} [delete]
func (h *OrderHandler) DeleteOrder(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid order ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := h.service.DeleteOrder(uint(id)); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete order", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Order deleted successfully", requestID))
}

// UpdateOrderStatus godoc
// @Summary Update order status
// @Description Update the status of an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param status body object{status="string"} true "Status data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/status [patch]
func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid order ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	var request struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := h.service.UpdateOrderStatus(uint(id), request.Status); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update order status", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Order status updated successfully", requestID))
}

// UpdatePaymentStatus godoc
// @Summary Update payment status
// @Description Update the payment status of an order by ID
// @Tags orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Param payment body object{status="string",payment_id="string"} true "Payment data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /orders/{id}/payment [patch]
func (h *OrderHandler) UpdatePaymentStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid order ID"},
		}
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	var request struct {
		Status    string  `json:"status" binding:"required"`
		PaymentID *string `json:"payment_id"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		c.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := h.service.UpdatePaymentStatus(uint(id), request.Status, request.PaymentID); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update payment status", err.Error(), requestID)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	c.JSON(http.StatusOK, utils.SuccessResponse(nil, "Payment status updated successfully", requestID))
}
