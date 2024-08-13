package ebook

import (
	"time"

	"gorm.io/gorm"
)

type Ebook struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	UserID int `gorm:"type:int(11)"`
	Ebook string `gorm:"type:text"`
	ThumbnailBook string `gorm:"type:text"`
	Genre string `gorm:"type:varchar(255)"`
	Title string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Author string `gorm:"type:varchar(255)"`
	BookCreated time.Time
}

