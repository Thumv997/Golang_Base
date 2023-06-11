package models

import (
	"time"

	"gorm.io/gorm"
)

// User represents a user in the application.
type User struct {
	gorm.Model
	ID        uint           `gorm:"primaryKey;autoIncrement"`
	Name      string         `gorm:"not null"`
	Email     string         `gorm:"unique;not null"`
	Password  string         `gorm:"not null"`
	Role      string		 `gorm:"not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
