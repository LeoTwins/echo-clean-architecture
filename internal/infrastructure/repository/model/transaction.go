package model

import "time"

type Transaction struct {
	ID        uint `gorm:"primaryKey"`
	AccountID uint
	Account   Account   `gorm:"foreignKey:AccountID"`
	Type      string    `gorm:"not null"`
	Amount    uint      `gorm:"not null"`
	Date      time.Time `gorm:"autoCreateTime"`
}
