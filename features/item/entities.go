package item

import (
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
}

