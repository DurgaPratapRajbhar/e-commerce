package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/services"

	"github.com/gin-gonic/gin"
)

type UserAddressController struct {
	service services.UserAddressService
}

func NewUserAddressController(service services.UserAddressService) *UserAddressController {
	return &UserAddressController{service: service}
}

func (c *UserAddressController) CreateAddress(ctx *gin.Context) {
	var address models.UserAddress
	if err := ctx.ShouldBindJSON(&address); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.CreateAddress(&address); err != nil {
		requestID := utils.GenerateRequestID()
		if strings.Contains(err.Error(), utils.ErrAlreadyExists) {
			
			validationErrors := []utils.ValidationError{
				{Field: "address", Message: "address already exists"},
			}
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(address, "Address created successfully", requestID))
}

func (c *UserAddressController) GetAddressByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid address ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	address, err := c.service.GetAddressByID(uint(id))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(address, "Address retrieved successfully", requestID))
}

func (c *UserAddressController) GetAddressesByUserID(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	addresses, err := c.service.GetAddressesByUserID(uint(userID))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get addresses", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(addresses, "Addresses retrieved successfully", requestID))
}

func (c *UserAddressController) GetAddressByUserIDAndType(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	addressType := ctx.Param("type")

	address, err := c.service.GetAddressByUserIDAndType(uint(userID), addressType)
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(address, "Address retrieved successfully", requestID))
}

func (c *UserAddressController) UpdateAddress(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid address ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	var address models.UserAddress
	if err := ctx.ShouldBindJSON(&address); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.UpdateAddress(uint(id), &address); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(address, "Address updated successfully", requestID))
}

func (c *UserAddressController) DeleteAddress(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid address ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.DeleteAddress(uint(id)); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(map[string]string{"message": "Address deleted successfully"}, "Address deleted successfully", requestID))
}

func (c *UserAddressController) SetDefaultAddress(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	addressID, err := strconv.ParseUint(ctx.Param("address_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "address_id", Message: "invalid address ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.SetDefaultAddress(uint(userID), uint(addressID)); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to set default address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(map[string]string{"message": "Default address set successfully"}, "Default address set successfully", requestID))
}

func (c *UserAddressController) GetDefaultAddress(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	address, err := c.service.GetDefaultAddress(uint(userID))
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get default address", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(address, "Default address retrieved successfully", requestID))
}