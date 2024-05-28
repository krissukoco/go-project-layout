package user_repository

import (
	"context"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
)

var (
	ErrNotFound = errors.New("not found")
)

type Repository interface {
	// Get gets user by id
	Get(ctx context.Context, id int64) (*entity.User, error)
	// GetByEmail gets user by email
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
}
