package auth

import (
	"context"
	"errors"
	desc "github.com/DenisCom3/m-auth/pkg/auth_v1"
	"time"
)

func (i *Implementation) Login(ctx context.Context, r *desc.LoginRequest) (*desc.LoginResponse, error) {

	user, err := i.userService.GetByName(ctx, r.GetUsername())

	if err != nil {
		return nil, err
	}

	// TODO: add password checking

	refreshToken, err := i.authService.GenerateToken(user.Info, []byte("lsjdjshfbdsjdhb"), 5*time.Minute)

	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &desc.LoginResponse{RefreshToken: refreshToken}, nil
}
