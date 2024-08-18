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
	GenreEbook string `gorm:"type:varchar(255)"`
	TitleEbook string `gorm:"type:varchar(255)"`
	DescriptionEbook string `gorm:"type:varchar(255)"`
	AuthorEbook string `gorm:"type:varchar(255)"`
	BookCreated time.Time
}

