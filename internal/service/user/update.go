package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
)

func (s *serv) Update(ctx context.Context, user model.UpdateUser) error {
	err := s.userRepo.Update(ctx, user)

	if err != nil {
		return err
	}

	return nil
}
