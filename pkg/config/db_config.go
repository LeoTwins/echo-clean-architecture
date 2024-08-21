package config

import (
	"os"
)

type DBConfig struct {
	Host     string
	Name     string
	Port     string
	User     string
	Passwrod string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Passwrod: os.Getenv("DB_PASSWORD"),
	}
}
