package kafka

import (
	"context"
	"github.com/IBM/sarama"
)

type ConsumerHandler func(ctx context.Context, msg *sarama.ConsumerMessage) error

type Consumer interface {
	Consume(ctx context.Context, topicName string, h ConsumerHandler) (err error)
	Close() error
}
