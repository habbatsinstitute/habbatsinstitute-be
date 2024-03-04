package realtimechat

import (
	"time"

	"gorm.io/gorm"
)

type PrivateChatHistory struct {
	SenderID  int `bson:"sender_id"`
	RecipientID int `bson:"recipient_id"`
	Message []Chat
}

type Chat struct {
	Text      string `bson:"text"`
	Blob      string `bson:"blob"`
	Timestamp time.Time `bson:"timestamp"`
}

type User struct {
	ID             int    		`gorm:"primaryKey;type:int(11)"`
	RoleID         int    		`gorm:"type:int(1);default:1"`
	Username       string 		`gorm:"type:varchar(255);not null"`
	Password       string 		`gorm:"type:varchar(255);not null"`
	ExpiryDate 	   time.Time
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}