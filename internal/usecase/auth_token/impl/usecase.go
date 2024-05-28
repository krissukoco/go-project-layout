package auth_token_usecase_impl

import (
	"context"

	"github.com/krissukoco/go-project-layout/internal/entity"
	user_repository "github.com/krissukoco/go-project-layout/internal/repository/user"
	auth_token_usecase "github.com/krissukoco/go-project-layout/internal/usecase/auth_token"
)

type usecase struct {
	jwtSecret string
	userRepo  user_repository.Repository
}

func New(
	jwtSecret string,
	userRepo user_repository.Repository,
) auth_token_usecase.Usecase {
	return &usecase{
		jwtSecret: jwtSecret,
		userRepo:  userRepo,
	}
}

// Validate validates access token. Returns `ErrInvalidToken` if token is expired or invalid
func (uc *usecase) Validate(ctx context.Context, accessToken string) (int64, error) {
	// Parse JWT with secret, get user id
	panic("unimplemented")
}

// Refresh takes access and refresh tokens, validates, and give new access and refresh tokens.
//
// Returns `ErrInvalidToken` if accessToken or refreshToken are invalid
// or `ErrRefreshTokenExpired` if refresh token is expired
func (uc *usecase) Refresh(ctx context.Context, accessToken string, refreshToken string) (*entity.AuthTokens, error) {
	// Validate tokens
	// Confirm if user still exists
	// Build new token
	panic("unimplemented")
}
