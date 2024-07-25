package user

import (
	"github.com/DenisCom3/m-auth/internal/service"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
