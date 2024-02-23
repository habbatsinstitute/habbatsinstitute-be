package realtime_chat

import (
	"gorm.io/gorm"
)

type Realtime_chat struct {
	gorm.Model

	ID int `gorm:"type:int(11)"`
	Name string `gorm:"type:varchar(255)"`
}

