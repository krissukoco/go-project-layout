package auth_user_usecase

import (
	"context"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type Usecase interface {
	// Login validates user credentials and returns auth tokens if valid.
	//
	// Returns `ErrInvalidCredentials` if email is not registered or email and password do not match
	Login(ctx context.Context, email, password string) (*entity.AuthTokens, error)
}
