package mysql

import (
	"context"
	"database/sql"
	"time"

	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
)

type verificationRepository struct {
	db *sql.DB
}

func NewVerificationRepository(db *sql.DB) *verificationRepository {
	return &verificationRepository{db: db}
}

func (r *verificationRepository) Create(ctx context.Context, token *entity.VerificationToken) error {
	query := `INSERT INTO verification_tokens (user_id, token, expires_at, created_at, used_at, is_used) 
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

func (r *verificationRepository) FindByToken(ctx context.Context, token string) (*entity.VerificationToken, error) {
	query := `SELECT id, user_id, token, expires_at, created_at, used_at, is_used 
			  FROM verification_tokens WHERE token = ?`

	verificationToken := &entity.VerificationToken{}
	err := r.db.QueryRowContext(ctx, query, token).Scan(
		&verificationToken.ID, &verificationToken.UserID, &verificationToken.Token,
		&verificationToken.ExpiresAt, &verificationToken.CreatedAt, &verificationToken.UsedAt,
		&verificationToken.IsUsed,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return verificationToken, nil
}

func (r *verificationRepository) MarkAsUsed(ctx context.Context, id int64) error {
	query := `UPDATE verification_tokens SET is_used = ?, used_at = ? WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, true, time.Now(), id)
	return err
}

func (r *verificationRepository) DeleteExpired(ctx context.Context) error {
	query := `DELETE FROM verification_tokens WHERE expires_at < ?`
	_, err := r.db.ExecContext(ctx, query, time.Now())
	return err
}