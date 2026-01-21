package usecase

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/service"
	"time"
)

type AuthUseCase struct {
	userRepo     repository.UserRepository
	tokenService service.TokenService
}

func NewAuthUseCase(userRepo repository.UserRepository, tokenService service.TokenService) *AuthUseCase {
	return &AuthUseCase{
		userRepo:     userRepo,
		tokenService: tokenService,
	}
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token     string        `json:"token"`
	User      *entity.User  `json:"user"`
	ExpiresIn time.Duration `json:"expires_in"`
}

func (uc *AuthUseCase) Register(ctx context.Context, req RegisterRequest) (*entity.User, error) {
	existingUser, _ := uc.userRepo.FindByEmail(ctx, req.Email)
	if existingUser != nil {
		return nil, utils.GetError(utils.ErrAlreadyExists)
	}
	
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	
	role := req.Role
	if role == "" {
		role = "user"
	}
	
	user := &entity.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: string(hashedPassword),
		Role:         role,
		IsActive:     true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	
	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}
	
	return user, nil
}

func (uc *AuthUseCase) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, utils.GetError(utils.ErrInvalidCredentials)
	}
	
	if !user.IsActive {
		return nil, utils.GetError(utils.ErrUnauthorized)
	}
	
	if !utils.CheckPassword(req.Password, user.PasswordHash) {
		return nil, utils.GetError(utils.ErrInvalidCredentials)
	}
	
	token, err := uc.tokenService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, err
	}
	
	return &LoginResponse{
		Token:     token,
		User:      user,
		ExpiresIn: uc.tokenService.GetTokenExpiry(),
	}, nil
}

func (uc *AuthUseCase) GetUserByID(ctx context.Context, userID int64) (*entity.User, error) {
	return uc.userRepo.FindByID(ctx, userID)
}

func (uc *AuthUseCase) GetAllUsers(ctx context.Context, offset int, limit int) ([]*entity.User, error) {
	return uc.userRepo.FindAll(ctx, offset, limit)
}

func (uc *AuthUseCase) GetTotalUserCount(ctx context.Context) (int, error) {
	return uc.userRepo.CountAll(ctx)
}

 