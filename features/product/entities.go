package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
	Images string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(255)"`
	Composition string `gorm:"type:varchar(255)"`
	PomTR string `gorm:"type:varchar(255)"`
	MarketPlace string `gorm:"type:varchar(255)"`
	Whatsapp string `gorm:"type:varchar(255)"`
	Price int64 `gorm:"type:int(15)"`
	Quantity int `gorm:"type:int(11)"`
	ProductCreated time.Time
}

