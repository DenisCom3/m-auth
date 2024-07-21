package app

import (
	"context"
	userApi "github.com/DenisCom3/m-auth/internal/api/user"
	"github.com/DenisCom3/m-auth/internal/client/db"
	"github.com/DenisCom3/m-auth/internal/client/db/postgres"
	"github.com/DenisCom3/m-auth/internal/client/db/transaction"
	"github.com/DenisCom3/m-auth/internal/config"
	"github.com/DenisCom3/m-auth/internal/repository"
	userRepo "github.com/DenisCom3/m-auth/internal/repository/user"
	"github.com/DenisCom3/m-auth/internal/service"
	userServ "github.com/DenisCom3/m-auth/internal/service/user"
	"log"
)

// in provider only interface or pointer to struct
type serviceProvider struct {
	pgConfig   config.Postgres
	grpcConfig config.Grpc

	dbClient  db.Client
	txManager db.TxManager

	userRepository repository.UserRepository
	userService    service.UserService

	userImpl *userApi.Implementation
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
	if s.userRepository != nil {
		s.userRepository = userRepo.New(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService != nil {
		s.userService = userServ.New(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *userApi.Implementation {
	if s.userImpl == nil {
		s.userImpl = userApi.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}
