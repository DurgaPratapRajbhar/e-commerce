package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type tokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *tokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(ctx context.Context, token *entity.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (user_id, token, expires_at, created_at, revoked_at, is_revoked) 
			  VALUES (?, ?, ?, ?, ?, ?)`

	result, err := r.db.ExecContext(ctx, query, token.UserID, token.Token, token.ExpiresAt, 
		token.CreatedAt, token.RevokedAt, token.IsRevoked)
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

func (r *tokenRepository) FindByToken(ctx context.Context, token string) (*entity.RefreshToken, error) {
	query := `SELECT id, user_id, token, expires_at, created_at, revoked_at, is_revoked 
			  FROM refresh_tokens WHERE token = ?`

	refreshToken := &entity.RefreshToken{}
	err := r.db.QueryRowContext(ctx, query, token).Scan(
		&refreshToken.ID, &refreshToken.UserID, &refreshToken.Token,
		&refreshToken.ExpiresAt, &refreshToken.CreatedAt, &refreshToken.RevokedAt,
		&refreshToken.IsRevoked,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return refreshToken, nil
}

func (r *tokenRepository) Revoke(ctx context.Context, id int64) error {
	query := `UPDATE refresh_tokens SET is_revoked = ?, revoked_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, true, time.Now(), id)
	return err
}

func (r *tokenRepository) RevokeAllForUser(ctx context.Context, userID int64) error {
	query := `UPDATE refresh_tokens SET is_revoked = ?, revoked_at = ? WHERE user_id = ? AND is_revoked = FALSE`
	_, err := r.db.ExecContext(ctx, query, true, time.Now(), userID)
	return err
}

func (r *tokenRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM refresh_tokens WHERE expires_at < ?`
	_, err := r.db.ExecContext(ctx, query, time.Now())
	return err
}