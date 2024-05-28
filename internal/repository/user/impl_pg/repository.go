package user_repository_impl_pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
	user_repository "github.com/krissukoco/go-project-layout/internal/repository/user"
)

type repository struct {
	db *sql.DB
}

var _ user_repository.Repository = (*repository)(nil)

func New(db *sql.DB) user_repository.Repository {
	return &repository{db}
}

func (r *repository) Get(ctx context.Context, id int64) (*entity.User, error) {
	q := `
		SELECT
			id, name, email, password, created_at, updated_at
		FROM users u
		WHERE u.id = $1
	`
	args := []interface{}{id}

	var x entity.User
	err := r.db.QueryRowContext(ctx, q, args...).Scan(
		&x.Id, &x.Name, &x.Email, &x.Password, &x.CreatedAt, &x.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_repository.ErrNotFound
		}
		return nil, err
	}
	return &x, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	q := `
		SELECT
			id, name, email, password, created_at, updated_at
		FROM users u
		WHERE u.email = $1
	`
	args := []interface{}{email}

	var x entity.User
	err := r.db.QueryRowContext(ctx, q, args...).Scan(
		&x.Id, &x.Name, &x.Email, &x.Password, &x.CreatedAt, &x.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_repository.ErrNotFound
		}
		return nil, err
	}
	return &x, nil
}
