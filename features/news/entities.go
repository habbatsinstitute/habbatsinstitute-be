package news

import (
	"time"

	"gorm.io/gorm"
)

type News struct {
	gorm.Model

	ID 			int 	`gorm:"type:int(11)"`
	UserID		int 	`gorm:"type:int(11)"`
	Category 	string	`gorm:"type:string"`
	Images		string	`gorm:"type:text"`
	Title		string	`gorm:"type:text"`
	Description	string	`gorm:"type:text"`
	Views		int		`gorm:"type:int(11)"`
	NewsCreated time.Time 
}

type Category struct {
	ID 			int `gorm:"type:int(11)"`
	UserID		int 	`gorm:"type:int(11)"`
	Name 		string `gorm:"type:varchar(255)"`
}

