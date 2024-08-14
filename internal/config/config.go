package config

import (
	"fmt"
	"github.com/DenisCom3/m-auth/internal/config/grpc"
	"github.com/DenisCom3/m-auth/internal/config/kafka"
	"github.com/DenisCom3/m-auth/internal/config/postgres"
	"github.com/DenisCom3/m-auth/internal/config/redis"
	"github.com/IBM/sarama"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	"time"
)

var cfg *config

type yamlConfig struct {
	Postgres postgres.Postgres `yaml:"postgres" env-required:"true"`
	Grpc     grpc.Grpc         `yaml:"grpc" env-required:"true"`
	Redis    redis.Redis       `yaml:"redis" env-required:"true"`
	Kafka    kafka.Kafka       `yaml:"kafka"`
}

type config struct {
	postgres      Postgres
	grpc          Grpc
	redis         Redis
	kafkaConsumer KafkaConsumer
}

type Postgres interface {
	Dsn() string
}

type Grpc interface {
	Address() string
}

type Redis interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

type KafkaConsumer interface {
	Brokers() []string
	GroupID() string
	Config() *sarama.Config
}

func GetPostgres() Postgres {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.postgres
}

func GetRedis() Redis {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.redis
}

func GetGrpc() Grpc {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.grpc
}

func GetKafkaConsumer() KafkaConsumer {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.kafkaConsumer
}
func MustLoad() error {

	if cfg != nil {
		return fmt.Errorf("config already initialized")
	}

	err := godotenv.Load()

	if err != nil {
		return fmt.Errorf("%w", err)
	}
	configPath := os.Getenv("APP_CONFIG_PATH")

	if configPath == "" {
		return fmt.Errorf("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist: %s", configPath)
	}

	var yaml yamlConfig

	if err := cleanenv.ReadConfig(configPath, &yaml); err != nil {
		return fmt.Errorf("cannot read config: %s", err)
	}

	cfg = &config{
		postgres:      yaml.Postgres,
		grpc:          yaml.Grpc,
		redis:         yaml.Redis,
		kafkaConsumer: yaml.Kafka.Consumer,
	}
	return nil
}
