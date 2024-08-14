package user

import (
	"github.com/DenisCom3/m-auth/internal/client/cache"
	"github.com/DenisCom3/m-auth/internal/repository"
	"github.com/DenisCom3/m-auth/internal/service"
	"github.com/DenisCom3/m-auth/internal/service/crypto"
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

func New(r repository.UserRepository, c cache.Cache) service.UserService {
	return &serv{
		cache:       c,
		userRepo:    r,
		hashService: crypto.Service{},
	}
}
