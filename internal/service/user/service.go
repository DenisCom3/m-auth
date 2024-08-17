package user

import (
	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/repository"
	"github.com/DenisCom3/m-auth/internal/service"
)

type hashService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type serv struct {
	cache cache.Cache

	userRepo    repository.UserRepository
	hashService hashService
}

func New(r repository.UserRepository, c cache.Cache, h hashService) service.UserService {
	return &serv{
		cache:       c,
		userRepo:    r,
		hashService: h,
	}
}
