package user_repository_mock

import (
	"context"

	"github.com/krissukoco/go-project-layout/internal/entity"
	user_repository "github.com/krissukoco/go-project-layout/internal/repository/user"
)

type mock struct {
	items []entity.User
}

var _ user_repository.Repository = (*mock)(nil)

func New(items []entity.User) user_repository.Repository {
	return &mock{items}
}

func (m *mock) Get(ctx context.Context, id int64) (*entity.User, error) {
	for _, v := range m.items {
		if v.Id == id {
			return &v, nil
		}
	}
	return nil, user_repository.ErrNotFound
}

func (m *mock) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	for _, v := range m.items {
		if v.Email == email {
			return &v, nil
		}
	}
	return nil, user_repository.ErrNotFound
}
