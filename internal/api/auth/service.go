package auth

import (
	"github.com/DenisCom3/m-auth/internal/service"
	desc "github.com/DenisCom3/m-auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
	userService service.UserService
}

func NewImplementation(authServ service.AuthService, userServ service.UserService) *Implementation {
	return &Implementation{
		authService: authServ,
		userService: userServ,
	}
}
