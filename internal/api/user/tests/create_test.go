package tests

import (
	"context"
	"fmt"
	api "github.com/DenisCom3/m-auth/internal/api/user"
	"github.com/DenisCom3/m-auth/internal/api/user/errors"
	"github.com/DenisCom3/m-auth/internal/model"
	"github.com/DenisCom3/m-auth/internal/service"
	servM "github.com/DenisCom3/m-auth/internal/service/mocks"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 8)
		role     = 0

		// request with password confirmation
		req1 = &desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: password,
			Role:            desc.Role(role),
		}

		serviceErr = fmt.Errorf("service error")

		// request without password confirmation
		req2 = desc.CreateRequest{
			Name:            name,
			Email:           email,
			Password:        password,
			PasswordConfirm: "not equals pass",
			Role:            desc.Role(role),
		}

		user = model.CreateUser{
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  model.Role(role),
			},
			Password: password,
		}

		res = &desc.CreateResponse{Id: id}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req1,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := servM.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(id, nil)

				return mock
			},
		},
		{
			name: "failed case (pass not eq)",
			args: args{
				ctx: ctx,
				req: &req2,
			},
			want: nil,
			err:  errors.ErrPassNotEq,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := servM.NewUserServiceMock(mc)

				return mock
			},
		},
		{
			name: "failed case (service err)",
			args: args{
				ctx: ctx,
				req: req1,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := servM.NewUserServiceMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(0, serviceErr)

				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServMock := tt.userServiceMock(mc)
			implementation := api.NewImplementation(userServMock)

			res, err := implementation.Create(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, res)

		})
	}

}
