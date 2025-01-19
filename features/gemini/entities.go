package gemini

import (
	"gorm.io/gorm"
)

type Information struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Question string `gorm:"type:varchar(255)"`
	Answer string `gorm:"type:text"`
}

