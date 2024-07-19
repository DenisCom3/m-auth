package user

import (
	"github.com/DenisCom3/m-auth/internal/repository"
	"github.com/DenisCom3/m-auth/internal/service"
	"github.com/DenisCom3/m-auth/internal/service/crypto"
)

type hashService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}

type serv struct {
	userRepo    repository.UserRepository
	hashService hashService
}

func New(r repository.UserRepository) service.UserService {
	return &serv{
		userRepo:    r,
		hashService: crypto.Service{},
	}
}
