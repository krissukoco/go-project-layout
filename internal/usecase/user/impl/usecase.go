package user_usecase_impl

import (
	"context"
	"errors"

	"github.com/krissukoco/go-project-layout/internal/entity"
	user_repository "github.com/krissukoco/go-project-layout/internal/repository/user"
	user_usecase "github.com/krissukoco/go-project-layout/internal/usecase/user"
)

type usecase struct {
	repo user_repository.Repository
}

func New(
	repo user_repository.Repository,
) user_usecase.Usecase {
	return &usecase{repo}
}

func (uc *usecase) GetProfile(ctx context.Context, id int64) (*entity.User, error) {
	user, err := uc.repo.Get(ctx, id)
	if err != nil {
		if errors.Is(err, user_repository.ErrNotFound) {
			return nil, user_usecase.ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
