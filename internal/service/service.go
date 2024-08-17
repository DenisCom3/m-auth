package service

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
	"time"
)

type AuthService interface {
	CheckPasswordHash(password, hash string) bool
	GenerateToken(info model.UserInfo, secretKey []byte, duration time.Duration) (string, error)
	VerifyToken(tokenStr string, secretKey []byte) (*model.UserClaims, error)
	HashPassword(password string) (string, error)
}

type UserService interface {
	GetById(ctx context.Context, id int64) (*model.User, error)
	GetByName(ctx context.Context, name string) (*model.User, error)
	Create(ctx context.Context, user model.CreateUser) (int64, error)
	Update(ctx context.Context, user model.UpdateUser) error
	Delete(ctx context.Context, id int64) error
}

type ConsumerService interface {
	RunConsumer(ctx context.Context) error
}
