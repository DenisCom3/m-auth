package repository

import (
	"context"
	"errors"

	"github.com/DenisCom3/m-auth/internal/model"
)

type UserRepository interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, user model.CreateUser) (int64, error)
	Update(ctx context.Context, user model.UpdateUser) error
	Delete(ctx context.Context, id int64) error
}

var (
	ErrUserNotFound = errors.New("user not found in db")
)
