package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	DBConfig DBConfig
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		DBConfig: LoadDBConfig(),
	}, nil
}
