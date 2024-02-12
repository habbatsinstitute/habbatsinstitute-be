package course

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model

	ID 				int 	`gorm:"type:int(11)"`
	UserID			int 	`gorm:"type:int(11)"`
	MediaFile		string	`gorm:"type:text"`
	Title			string	`gorm:"type:text"`
	Description		string	`gorm:"type:text"`
	Author			string	`gorm:"type:varchar(255)"`
	Views 			int     `gorm:"type:int(11)"`
	CourseCreatedAt	time.Time 
}

