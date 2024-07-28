package tests

import (
	"context"
	"database/sql"
	"fmt"
	api "github.com/DenisCom3/m-auth/internal/api/user"
	"github.com/DenisCom3/m-auth/internal/model"
	"github.com/DenisCom3/m-auth/internal/service"
	servM "github.com/DenisCom3/m-auth/internal/service/mocks"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	t.Parallel()

	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id    = gofakeit.Int64()
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = 0
		now   = time.Now()

		req = &desc.GetRequest{Id: id}

		serviceErr = fmt.Errorf("service err")

		user = &model.User{
			ID: id,
			Info: model.UserInfo{
				Name:  name,
				Email: email,
				Role:  model.Role(role),
			},
			CreatedAt: now,
			UpdatedAt: sql.NullTime{Time: now, Valid: true},
		}

		res = &desc.GetResponse{Id: id,
			Name:      name,
			Email:     email,
			Role:      desc.Role(role),
			CreatedAt: timestamppb.New(now),
			UpdatedAt: timestamppb.New(now),
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := servM.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(user, nil)

				return mock

			},
		},
		{
			name: "failed case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := servM.NewUserServiceMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)

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

			res, err := implementation.Get(tt.args.ctx, tt.args.req)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, res)

		})
	}

}
