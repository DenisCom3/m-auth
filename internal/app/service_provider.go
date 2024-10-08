package app

import (
	"context"
	"github.com/DenisCom3/m-auth/internal/client/broker_message/kafka"
	"github.com/DenisCom3/m-auth/internal/client/broker_message/kafka/consumer"
	"github.com/DenisCom3/m-auth/internal/closer"
	"github.com/DenisCom3/m-auth/internal/service/consumer/user_creater"
	"github.com/IBM/sarama"
	"log"

	authApi "github.com/DenisCom3/m-auth/internal/api/auth"
	userApi "github.com/DenisCom3/m-auth/internal/api/user"
	"github.com/DenisCom3/m-auth/internal/client/cache"
	goredis "github.com/DenisCom3/m-auth/internal/client/cache/go-redis"
	"github.com/DenisCom3/m-auth/internal/client/db"
	"github.com/DenisCom3/m-auth/internal/client/db/postgres"
	"github.com/DenisCom3/m-auth/internal/client/db/transaction"
	"github.com/DenisCom3/m-auth/internal/config"
	"github.com/DenisCom3/m-auth/internal/repository"
	userRepo "github.com/DenisCom3/m-auth/internal/repository/user"
	"github.com/DenisCom3/m-auth/internal/service"
	authServ "github.com/DenisCom3/m-auth/internal/service/auth"
	userServ "github.com/DenisCom3/m-auth/internal/service/user"
)

// in provider only interface or pointer to struct
type serviceProvider struct {
	pgConfig            config.Postgres
	grpcConfig          config.Grpc
	redisConfig         config.Redis
	kafkaConsumerConfig config.KafkaConsumer
	cacheClient         cache.Cache
	dbClient            db.Client
	txManager           db.TxManager

	consumer            kafka.Consumer
	userCreaterConsumer service.ConsumerService
	userRepository      repository.UserRepository
	userService         service.UserService
	authService         service.AuthService

	userImpl *userApi.Implementation
	authImpl *authApi.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}
func (s *serviceProvider) GRPCConfig() config.Grpc {
	if s.grpcConfig == nil {
		cfg := config.GetGrpc()
		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) PostgresConfig() config.Postgres {
	if s.pgConfig == nil {
		cfg := config.GetPostgres()

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) RedisConfig() config.Redis {
	if s.redisConfig == nil {
		cfg := config.GetRedis()

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) KafkaConsumerConfig() config.KafkaConsumer {
	if s.kafkaConsumerConfig == nil {
		cfg := config.GetKafkaConsumer()

		s.kafkaConsumerConfig = cfg
	}

	return s.kafkaConsumerConfig
}

func (s *serviceProvider) CacheClient(_ context.Context) cache.Cache {
	if s.cacheClient == nil {
		s.cacheClient = goredis.NewCache(s.RedisConfig())
	}

	return s.cacheClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		client, err := postgres.New(ctx, s.PostgresConfig().Dsn())

		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = client.DB().Ping(ctx)

		if err != nil {
			log.Fatalf("ping error: %v", err)
		}

		s.dbClient = client
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.New(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepo.New(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userServ.New(s.UserRepository(ctx), s.CacheClient(ctx), s.AuthService(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *userApi.Implementation {
	if s.userImpl == nil {
		s.userImpl = userApi.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
func (s *serviceProvider) AuthImpl(ctx context.Context) *authApi.Implementation {
	if s.authImpl == nil {
		s.authImpl = authApi.NewImplementation(s.AuthService(ctx), s.UserService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) KafkaConsumer(_ context.Context) kafka.Consumer {
	if s.consumer == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			s.KafkaConsumerConfig().Brokers(),
			s.KafkaConsumerConfig().GroupID(),
			s.KafkaConsumerConfig().Config(),
		)

		if err != nil {
			log.Fatalf("failed to create consumer group: %v", err)
		}

		consumerGroupHandler := consumer.NewGroupHandler()
		s.consumer = consumer.New(consumerGroup, consumerGroupHandler)
		closer.Add(s.consumer.Close)
	}

	return s.consumer
}

func (s *serviceProvider) UserCreaterConsumer(ctx context.Context) service.ConsumerService {
	if s.userCreaterConsumer == nil {
		s.userCreaterConsumer = user_creater.New(s.UserService(ctx), s.KafkaConsumer(ctx))
	}

	return s.userCreaterConsumer
}

func (s *serviceProvider) AuthService(_ context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authServ.New()
	}
	return s.authService
}
