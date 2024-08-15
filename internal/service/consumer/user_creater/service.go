package user_creater

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/client/broker_message/kafka"
	"github.com/DenisCom3/m-auth/internal/service"
)

const topic = "user-created"

type Service struct {
	userService service.UserService
	consumer    kafka.Consumer
}

func New(
	s service.UserService,
	c kafka.Consumer,
) *Service {
	return &Service{
		userService: s,
		consumer:    c,
	}
}

func (s *Service) RunConsumer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-s.run(ctx):
			if err != nil {
				return err
			}
		}
	}
}

func (s *Service) run(ctx context.Context) <-chan error {
	errChan := make(chan error)

	go func() {
		defer close(errChan)

		errChan <- s.consumer.Consume(ctx, topic, s.UserCreateHandler)
	}()

	return errChan
}
