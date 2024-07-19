package repository

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
)

type UserRepository interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user model.CreateUser) (int64, error)
}
