package user

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/converter"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
)

func (i *Implementation) Get(ctx context.Context, r *desc.GetRequest) (*desc.GetResponse, error) {

	user, err := i.userService.GetById(ctx, r.GetId())

	if err != nil {
		return nil, err
	}
	return converter.FromServiceToTransport(user), nil
}
