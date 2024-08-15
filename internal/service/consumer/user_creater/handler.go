package user_creater

import (
	"context"
	"encoding/json"
	"github.com/DenisCom3/m-auth/internal/model"
	"github.com/IBM/sarama"
)

func (s *Service) UserCreateHandler(ctx context.Context, msg *sarama.ConsumerMessage) error {
	createUser := &model.CreateUser{}
	err := json.Unmarshal(msg.Value, createUser)
	if err != nil {
		return err
	}

	_, err = s.userService.Create(ctx, *createUser)

	return err
}
