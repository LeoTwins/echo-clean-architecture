package db

import (
	"fmt"

	"github.com/LeoTwins/go-clean-architecture/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB(cfg config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Passwrod,
		cfg.Name,
		cfg.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
