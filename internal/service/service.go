package service

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
)

type UserService interface {
	Get(ctx context.Context, id int64) (model.User, error)
	Create(ctx context.Context, user model.CreateUser) (int64, error)
}
