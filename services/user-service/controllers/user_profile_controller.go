package controllers

import (
	"net/http"
	"strconv"
	"fmt"
	"strings"
 
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/models"
	"github.com/DurgaPratapRajbhar/e-commerce/user-service/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserProfileController struct {
	service services.UserProfileService
}

func NewUserProfileController(service services.UserProfileService) *UserProfileController {
	return &UserProfileController{service: service}
}

func (c *UserProfileController) CreateProfile(ctx *gin.Context) {
	var profile models.UserProfile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.CreateProfile(&profile); err != nil {
		requestID := utils.GenerateRequestID()
		if strings.Contains(err.Error(), utils.ErrAlreadyExists) {
			
			validationErrors := []utils.ValidationError{
				{Field: "user_id", Message: "profile already exists for this user"},
			}
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to create profile", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(profile, "Profile created successfully", requestID))
}

func (c *UserProfileController) GetProfileByUserID(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	fmt.Println("User ID:", userID)

	profile, err := c.service.GetProfileByUserID(uint(userID))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrNotFound, "User profile not found", "User profile not found", requestID)
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get profile", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(profile, "Profile retrieved successfully", requestID))
}

func (c *UserProfileController) GetProfileByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "id", Message: "invalid profile ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	profile, err := c.service.GetProfileByID(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			requestID := utils.GenerateRequestID()
			response := utils.ErrorResponse(utils.ErrNotFound, "User profile not found", "User profile not found", requestID)
			ctx.JSON(http.StatusNotFound, response)
			return
		}
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get profile", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(profile, "Profile retrieved successfully", requestID))
}

func (c *UserProfileController) UpdateProfile(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	var profile models.UserProfile
	if err := ctx.ShouldBindJSON(&profile); err != nil {
		validationErrors := utils.ParseValidationErrors(err.Error())
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.UpdateProfile(uint(userID), &profile); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to update profile", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(profile, "Profile updated successfully", requestID))
}

func (c *UserProfileController) DeleteProfile(ctx *gin.Context) {
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 32)
	if err != nil {
		validationErrors := []utils.ValidationError{
			{Field: "user_id", Message: "invalid user ID"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	if err := c.service.DeleteProfile(uint(userID)); err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to delete profile", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(map[string]string{"message": "Profile deleted successfully"}, "Profile deleted successfully", requestID))
}

func (c *UserProfileController) GetProfilesBulk(ctx *gin.Context) {
	userIDsStr := ctx.Query("user_ids")
	if userIDsStr == "" {
		validationErrors := []utils.ValidationError{
			{Field: "user_ids", Message: "user_ids parameter is required"},
		}
		requestID := utils.GenerateRequestID()
		ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
		return
	}

	userIDs := []uint{}
	idStrs := strings.Split(userIDsStr, ",")
	for _, idStr := range idStrs {
		id, err := strconv.ParseUint(strings.TrimSpace(idStr), 10, 32)
		if err != nil {
			validationErrors := []utils.ValidationError{
				{Field: "user_ids", Message: fmt.Sprintf("invalid user ID: %s", idStr)},
			}
			requestID := utils.GenerateRequestID()
			ctx.JSON(http.StatusBadRequest, utils.ValidationErrorResponse(validationErrors, requestID))
			return
		}
		userIDs = append(userIDs, uint(id))
	}

	profiles, err := c.service.GetProfilesByUserIDs(userIDs)
	if err != nil {
		requestID := utils.GenerateRequestID()
		response := utils.ErrorResponse(utils.ErrInternalServer, "Failed to get profiles", err.Error(), requestID)
		ctx.JSON(http.StatusInternalServerError, response)
		return
	}

	requestID := utils.GenerateRequestID()
	ctx.JSON(http.StatusOK, utils.SuccessResponse(profiles, "Profiles retrieved successfully", requestID))
}