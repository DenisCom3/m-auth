package converter

import (
	"github.com/DenisCom3/m-auth/internal/model"
	modelRepo "github.com/DenisCom3/m-auth/internal/repository/user/model"
)

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		ID: user.ID,
		Info: model.UserInfo{
			Name:  user.Name,
			Email: user.Email,
			Role:  model.Role(user.Role),
		},
	}
}
