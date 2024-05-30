package configs

import (
	"os"

	"github.com/joho/godotenv"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
)

type Config struct {
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	APIPort          string
	DBPort           string
}

func LoadEnv() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, errs.ErrLoadEnvVars
	}

	config := Config{
		DatabaseHost:     os.Getenv("DATABASE_HOST"),
		DatabaseUser:     os.Getenv("DATABASE_USER"),
		DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
		DatabaseName:     os.Getenv("DATABASE_NAME"),
		APIPort:          os.Getenv("API_PORT"),
		DBPort:           os.Getenv("DB_PORT"),
	}

	return &config, nil
}
