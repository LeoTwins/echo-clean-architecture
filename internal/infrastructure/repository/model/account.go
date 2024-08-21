package model

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey autoIncrement"`
	Name      string    `gorm:"not null"`
	Balance   uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
