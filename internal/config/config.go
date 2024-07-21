package config

import (
	"fmt"
	"github.com/DenisCom3/m-auth/internal/config/grpc"
	"github.com/DenisCom3/m-auth/internal/config/postgres"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
)

var cfg *Config

type yamlConfig struct {
	Postgres postgres.Postgres `yaml:"postgres" env-required:"true"`
	Grpc     grpc.Grpc         `yaml:"grpc" env-required:"true"`
}

type Config struct {
	postgres Postgres
	grpc     Grpc
}

type Postgres interface {
	Dsn() string
}

type Grpc interface {
	Address() string
}

func GetPostgres() Postgres {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.postgres
}

func GetGrpc() Grpc {
	if cfg == nil {
		panic("config not initialized")
	}
	return cfg.grpc
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

	cfg = &Config{
		postgres: yaml.Postgres,
		grpc:     yaml.Grpc,
	}
	return nil
}
