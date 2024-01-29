package user

import (
	"institute/features/course"
	"institute/features/news"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             int    		`gorm:"primaryKey;type:int(11)"`
	RoleID         int    		`gorm:"type:int(1);default:1"`
	Username       string 		`gorm:"type:varchar(255);not null"`
	Password       string 		`gorm:"type:varchar(255);not null"`
	ExpiryDate 	   time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`

	Course 		[]course.Course 	`gorm:"foreignKey:UserID;references:ID"`
	News 		[]news.News		 	`gorm:"foreignKey:UserID;references:ID"`
	Category 	[]news.Category 	`gorm:"foreignKey:UserID;references:ID"`
}

