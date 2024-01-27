package news

import (
	"gorm.io/gorm"
)

type News struct {
	gorm.Model

	ID 			int 	`gorm:"type:int(11)"`
	Category 	string	`gorm:"type:string"`
	Images		string	`gorm:"type:text"`
	Title		string	`gorm:"type:text"`
	Description	string	`gorm:"type:text"`
}

