package converter

import (
	"github.com/DenisCom3/m-auth/internal/model"
	desc "github.com/DenisCom3/m-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func FromServiceToTransport(user *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp

	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}
	return &desc.GetResponse{
		Id:        user.ID,
		Name:      user.Info.Name,
		Email:     user.Info.Email,
		Role:      desc.Role(user.Info.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
