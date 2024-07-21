package user

import (
	"github.com/DenisCom3/m-auth/internal/service"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	noteService service.UserService
}

func NewImplementation(noteService service.UserService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
