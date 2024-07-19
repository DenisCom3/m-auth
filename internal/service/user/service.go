package user

import (
	"github.com/DenisCom3/m-auth/internal/repository"
	"github.com/DenisCom3/m-auth/internal/service"
)

type serv struct {
	userRepo repository.UserRepository
}

func New(r repository.UserRepository) service.UserService {
	return &serv{userRepo: r}
}
