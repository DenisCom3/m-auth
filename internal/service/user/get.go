package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (model.User, error) {
	return model.User{}, nil
}
