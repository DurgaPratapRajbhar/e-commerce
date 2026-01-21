package repository

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type TokenRepository interface {
	Create(ctx context.Context, token *entity.RefreshToken) error
	FindByToken(ctx context.Context, token string) (*entity.RefreshToken, error)
	Revoke(ctx context.Context, id int64) error
	RevokeAllForUser(ctx context.Context, userID int64) error
	DeleteExpired(ctx context.Context) error
}