package auth_user_usecase_impl

import (
	"context"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
	user_repository "github.com/krissukoco/go-project-layout/internal/repository/user"
	auth_user_usecase "github.com/krissukoco/go-project-layout/internal/usecase/auth_user"
)

type usecase struct {
	userRepo user_repository.Repository
}

func New(
	userRepo user_repository.Repository,
) auth_user_usecase.Usecase {
	return &usecase{userRepo}
}

func (uc *usecase) validatePassword(input string, hash string) bool {
	// Use bcrypt etc.
	return true
}

func (uc *usecase) buildAccessToken(userId int64) string {
	// JWT stuff, put userId into sub as string
	return ""
}

func (uc *usecase) buildRefreshToken(userId int64) string {
	// JWT stuff, put userId into sub as string
	return ""
}

func (uc *usecase) Login(ctx context.Context, email, password string) (*entity.AuthTokens, error) {
	// Get user
	user, err := uc.userRepo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, user_repository.ErrNotFound) {
			return nil, auth_user_usecase.ErrInvalidCredentials
		}
		return nil, err
	}

	// Compare password
	if !uc.validatePassword(password, user.Password) {
		return nil, auth_user_usecase.ErrInvalidCredentials
	}

	accessToken := uc.buildAccessToken(user.Id)
	refreshToken := uc.buildRefreshToken(user.Id)

	return &entity.AuthTokens{accessToken, refreshToken}, nil
}
