package auth_token_usecase

import (
	"context"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
)

var (
	ErrRefreshTokenExpired = errors.New("refresh token is expired")
	ErrInvalidToken        = errors.New("token(s) are invalid")
)

type Usecase interface {
	// Validate validates access token. Returns `ErrInvalidToken` if token is expired or invalid
	Validate(ctx context.Context, accessToken string) (int64, error)
	// Refresh takes access and refresh tokens, validates, and give new access and refresh tokens.
	//
	// Returns `ErrInvalidToken` if accessToken or refreshToken are invalid
	// or `ErrRefreshTokenExpired` if refresh token is expired
	Refresh(ctx context.Context, accessToken string, refreshToken string) (*entity.AuthTokens, error)
}
