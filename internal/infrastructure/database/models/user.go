package models

import "time"

type UserModel struct {
	ID        uint64 `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex;not null"`
	Firstname string `gorm:"not null"`
	Lastname  string `gorm:"not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
