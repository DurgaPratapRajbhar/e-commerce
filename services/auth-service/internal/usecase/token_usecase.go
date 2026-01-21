package usecase

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/repository"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/service"
)

type TokenUseCase struct {
	tokenService service.TokenService
	userRepo     repository.UserRepository
}

func NewTokenUseCase(tokenService service.TokenService, userRepo repository.UserRepository) *TokenUseCase {
	return &TokenUseCase{
		tokenService: tokenService,
		userRepo:     userRepo,
	}
}

func (uc *TokenUseCase) ValidateToken(ctx context.Context, token string) (*service.TokenClaims, error) {
	claims, err := uc.tokenService.ValidateToken(token)
	if err != nil {
		return nil, err
	}
	
	user, err := uc.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return nil, utils.GetError(utils.ErrNotFound)
	}
	
	if !user.IsActive {
		return nil, utils.GetError(utils.ErrUnauthorized)
	}
	
	return claims, nil
}

func (uc *TokenUseCase) RefreshToken(ctx context.Context, token string) (string, error) {
	claims, err := uc.tokenService.ValidateToken(token)
	if err != nil {
		return "", err
	}
	
	user, err := uc.userRepo.FindByID(ctx, claims.UserID)
	if err != nil {
		return "", utils.GetError(utils.ErrNotFound)
	}
	
	if !user.IsActive {
		return "", utils.GetError(utils.ErrUnauthorized)
	}
	
	newToken, err := uc.tokenService.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}
	
	return newToken, nil
}