package auth

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int    `gorm:"primaryKey;type:int(11)"`
	RoleID         int    `gorm:"type:int(1);default:1"`
	Username       string `gorm:"type:varchar(255);not null"`
	Password       string `gorm:"type:varchar(255);not null"`
	ExpiryDate 	   time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}