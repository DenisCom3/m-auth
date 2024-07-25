package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
	"github.com/pkg/errors"
)

func (i *Implementation) Create(ctx context.Context, r *desc.CreateRequest) (*desc.CreateResponse, error) {

	if r.GetPassword() != r.GetPasswordConfirm() {
		return nil, errors.New("passwords not equals")
	}

	user := model.CreateUser{
		Info: model.UserInfo{
			Name:  r.GetName(),
			Email: r.GetEmail(),
			Role:  model.Role(r.GetRole()),
		},
		Password: r.GetPassword(),
	}
	id, err := i.userService.Create(ctx, user)

	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{Id: id}, nil
}
