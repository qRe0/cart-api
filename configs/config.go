package configs

import (
	"os"

	"github.com/joho/godotenv"
	errs "github.com/qRe0/innowise-cart-api/internal/errors"
)

type APIConfig struct {
	APIPort string
}

type DBConfig struct {
	DatabaseHost     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
	DBPort           string
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
			APIPort: os.Getenv("API_PORT"),
		},
		DB: DBConfig{
			DatabaseHost:     os.Getenv("DATABASE_HOST"),
			DatabaseUser:     os.Getenv("DATABASE_USER"),
			DatabasePassword: os.Getenv("DATABASE_PASSWORD"),
			DatabaseName:     os.Getenv("DATABASE_NAME"),
			DBPort:           os.Getenv("DB_PORT"),
		},
	}

	return &config, nil
}
