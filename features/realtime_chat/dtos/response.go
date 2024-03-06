package dtos

import (
	"time"
)

type RoomRes struct {
	RoomId int `json:"room_id"`
	SenderName string `json:"sender_name"`
	SenderId int `json:"sender_id"`
}

type ChatRes struct {
	SenderID    int       `json:"sender_id"`
	Text      string    `json:"text"`
	Blob      string    `json:"blob"`
	Timestamp time.Time `json:"timestamp"`
}