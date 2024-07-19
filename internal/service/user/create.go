package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
)

func (s *serv) Create(ctx context.Context, user model.CreateUser) (int64, error) {

	hashedPassword, err := s.hashService.HashPassword(user.Password)

	if err != nil {
		return 0, err
	}
	user.Password = hashedPassword

	id, err := s.userRepo.Create(ctx, user)

	if err != nil {
		return 0, err
	}
	return id, nil
}
