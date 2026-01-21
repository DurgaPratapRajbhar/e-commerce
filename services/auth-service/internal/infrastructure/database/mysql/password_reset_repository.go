package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type passwordResetRepository struct {
	db *sql.DB
}

func NewPasswordResetRepository(db *sql.DB) *passwordResetRepository {
	return &passwordResetRepository{db: db}
}

func (r *passwordResetRepository) Create(ctx context.Context, token *entity.PasswordResetToken) error {
	query := `INSERT INTO password_reset_tokens (user_id, token, expires_at, created_at, used_at, is_used) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, token.UserID, token.Token, token.ExpiresAt, 
		token.CreatedAt, token.UsedAt, token.IsUsed)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	token.ID = id
	return nil
}

func (r *passwordResetRepository) FindByToken(ctx context.Context, token string) (*entity.PasswordResetToken, error) {
	query := `SELECT id, user_id, token, expires_at, created_at, used_at, is_used 
			  FROM password_reset_tokens WHERE token = ?`

	passwordResetToken := &entity.PasswordResetToken{}
	err := r.db.QueryRowContext(ctx, query, token).Scan(
		&passwordResetToken.ID, &passwordResetToken.UserID, &passwordResetToken.Token,
		&passwordResetToken.ExpiresAt, &passwordResetToken.CreatedAt, &passwordResetToken.UsedAt,
		&passwordResetToken.IsUsed,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return passwordResetToken, nil
}

func (r *passwordResetRepository) MarkAsUsed(ctx context.Context, id int64) error {
	query := `UPDATE password_reset_tokens SET is_used = ?, used_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, true, time.Now(), id)
	return err
}

func (r *passwordResetRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM password_reset_tokens WHERE expires_at < ?`
	_, err := r.db.ExecContext(ctx, query, time.Now())
	return err
}