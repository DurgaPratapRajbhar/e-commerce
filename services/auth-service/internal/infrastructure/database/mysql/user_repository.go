package mysql

import (
	"context"
	"database/sql"
	"github.com/DurgaPratapRajbhar/e-commerce/auth-service/internal/domain/entity"
	"github.com/DurgaPratapRajbhar/e-commerce/pkg/utils"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (email, username, password_hash, role, is_active, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?)`
	
	result, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.PasswordHash, 
		user.Role, user.IsActive, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	
	user.ID = id
	return nil
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `SELECT id, email, username, password_hash, role, is_active, created_at, updated_at 
			  FROM users WHERE email = ?`
	
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash,
		&user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, utils.GetError(utils.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `SELECT id, email, username, password_hash, role, is_active, created_at, updated_at 
			  FROM users WHERE id = ?`
	
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash,
		&user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, utils.GetError(utils.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entity.User, error) {
	query := `SELECT id, email, username, password_hash, role, is_active, created_at, updated_at 
			  FROM users WHERE username = ?`
	
	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID, &user.Email, &user.Username, &user.PasswordHash,
		&user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
	)
	
	if err == sql.ErrNoRows {
		return nil, utils.GetError(utils.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}
	
	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *entity.User) error {
	query := `UPDATE users SET email = ?, username = ?, password_hash = ?, role = ?, 
			  is_active = ?, updated_at = ? WHERE id = ?`
	
	_, err := r.db.ExecContext(ctx, query, user.Email, user.Username, user.PasswordHash,
		user.Role, user.IsActive, user.UpdatedAt, user.ID)
	
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepository) FindAll(ctx context.Context, offset int, limit int) ([]*entity.User, error) {
	query := `SELECT id, email, username, password_hash, role, is_active, created_at, updated_at 
		  FROM users ORDER BY created_at DESC LIMIT ? OFFSET ?`
	
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	users := []*entity.User{}
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(
			&user.ID, &user.Email, &user.Username, &user.PasswordHash,
			&user.Role, &user.IsActive, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	
	return users, nil
}

func (r *userRepository) CountAll(ctx context.Context) (int, error) {
	query := `SELECT COUNT(*) FROM users`
	
	var count int
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	if err != nil {
		return 0, err
	}
	
	return count, nil
}

