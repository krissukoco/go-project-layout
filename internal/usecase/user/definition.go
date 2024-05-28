package user_usecase

import (
	"context"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type Usecase interface {
	GetProfile(ctx context.Context, id int64) (*entity.User, error)
}
