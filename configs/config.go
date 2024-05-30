package configs

import (
	"os"

	"github.com/joho/godotenv"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
)

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	User     string
	Password string
	Name     string
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

	config := Config{
		API: APIConfig{
			Port: os.Getenv("API_PORT"),
		},
		DB: DBConfig{
			Host:     os.Getenv("DATABASE_HOST"),
			User:     os.Getenv("DATABASE_USER"),
			Password: os.Getenv("DATABASE_PASSWORD"),
			Name:     os.Getenv("DATABASE_NAME"),
			Port:     os.Getenv("DB_PORT"),
		},
	}

	return &config, nil
}
