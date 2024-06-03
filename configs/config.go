package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
)

type APIConfig struct {
	Port            string
	ShutdownTimeout string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

type Config struct {
	API APIConfig
	DB  DBConfig
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errs.ErrLoadEnvVars
	}

	requiredEnvs := []string{
		"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT", "API_PORT", "SHUTDOWN_TIMEOUT",
	}

	for _, env := range requiredEnvs {
		if os.Getenv(env) == "" {
			return nil, fmt.Errorf("environment variable `%s` is not set or is empty", env)
		}
	}

	config := Config{
		API: APIConfig{
			Port:            os.Getenv("API_PORT"),
			ShutdownTimeout: os.Getenv("SHUTDOWN_TIMEOUT"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName:   os.Getenv("DB_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
	}

	return &config, nil
}
