package configs

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	errs "github.com/qRe0/cart-api/internal/errors"
)

type APIConfig struct {
	Port            string `env:"API_PORT"`
	ShutdownTimeout string `env:"SHUTDOWN_TIMEOUT"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	Port     string `env:"DB_PORT"`
}

type GRPC struct {
	Host string `env:"GRPC_HOST"`
	Port string `env:"GRPC_PORT"`
}

type Config struct {
	API  APIConfig
	DB   DBConfig
	GRPC GRPC
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errs.ErrLoadEnvVars
	}

	var apiCfg APIConfig
	err = env.Parse(&apiCfg)
	if err != nil {
		return nil, errors.Wrap(errs.ErrLoadEnvVars, "API")
	}

	var dbCfg DBConfig
	err = env.Parse(&dbCfg)
	if err != nil {
		return nil, errors.Wrap(errs.ErrLoadEnvVars, "DB")
	}

	var grpcCfg GRPC
	err = env.Parse(&grpcCfg)
	if err != nil {
		return nil, errors.Wrap(errs.ErrLoadEnvVars, "GRPC")
	}

	cfg := &Config{
		API:  apiCfg,
		DB:   dbCfg,
		GRPC: grpcCfg,
	}

	return cfg, nil
}
