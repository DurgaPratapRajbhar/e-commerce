package repository

import (
	"context"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id int64) error
	FindAll(ctx context.Context, offset int, limit int) ([]*entity.User, error)
	CountAll(ctx context.Context) (int, error)
}
