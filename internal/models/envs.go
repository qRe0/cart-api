package models

import "time"

type Config struct {
	DBHost          string        `env:"DB_HOST"`
	DBPort          int           `env:"DB_PORT"`
	DBUser          string        `env:"DB_USER"`
	DBPassword      string        `env:"DB_PASSWORD"`
	DBName          string        `env:"DB_NAME"`
	APIPort         int           `env:"API_PORT"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT"`
	GRPCHost        string        `env:"GRPC_HOST"`
	GRPCPort        int           `env:"GRPC_PORT"`
}
