package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/model"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
	"github.com/golang/protobuf/ptypes/empty"
)

func (i *Implementation) Update(ctx context.Context, r *desc.UpdateRequest) (*empty.Empty, error) {
	user := model.UpdateUser{
		ID: r.GetId(),
		Info: model.UserInfo{
			Name:  r.Name.Value,
			Email: r.Email.Value,
			Role:  model.Role(r.Role.Number()),
		},
	}
	err := i.userService.Update(ctx, user)

	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}
