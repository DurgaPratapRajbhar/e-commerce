package controllers

import (
	"net/http"
	"strconv"

	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/cart-service/services"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/gin-gonic/gin"
)

type CartController struct {
	cartService services.CartService
}

func NewCartController(service services.CartService) *CartController {
	return &CartController{
		cartService: service,
	}
}

func (ctrl *CartController) AddToCart(c *gin.Context) {
	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := ctrl.cartService.AddToCart(&cart); err != nil {
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to add to cart", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, utils.SuccessResponse(cart, "Item added to cart successfully", utils.GenerateRequestID()))
}

func (ctrl *CartController) GetCartByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", "Invalid ID", utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cart, err := ctrl.cartService.GetCartByID(id)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrNotFound, "Cart not found", "Cart not found", utils.GenerateRequestID())
		c.JSON(http.StatusNotFound, response)
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(cart, "Cart retrieved successfully", utils.GenerateRequestID()))
}

func (ctrl *CartController) GetCartByUserID(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", "Invalid User ID", utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	carts, err := ctrl.cartService.GetCartByUserID(userID)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to retrieve cart", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(carts, "Cart items retrieved successfully", utils.GenerateRequestID()))
}

func (ctrl *CartController) UpdateCart(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", "Invalid ID", utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var cart models.Cart
	if err := c.ShouldBindJSON(&cart); err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	cart.ID = id
	if err := ctrl.cartService.UpdateCart(&cart); err != nil {
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update cart", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(cart, "Cart updated successfully", utils.GenerateRequestID()))
}

func (ctrl *CartController) RemoveFromCart(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", "Invalid ID", utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := ctrl.cartService.RemoveFromCart(id); err != nil {
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to remove from cart", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.Status(http.StatusNoContent)
}

func (ctrl *CartController) ClearCart(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("userId"), 10, 64)
	if err != nil {
		response := utils.ErrorResponse(utils.ErrValidationFailed, "Invalid input", "Invalid User ID", utils.GenerateRequestID())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err := ctrl.cartService.ClearCart(userID); err != nil {
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to clear cart", err.Error(), utils.GenerateRequestID())
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	c.Status(http.StatusNoContent)
}