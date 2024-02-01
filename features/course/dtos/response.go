package dtos

import "time"

type ResCourse struct {
	ID              int       `json:"id" form:"id"`
	UserID          int       `json:"user_id" form:"user_id"`
	MediaFile       string    `json:"media_file" form:"media_file"`
	Title           string    `json:"title" form:"title"`
	Description     string    `json:"description" form:"description"`
	Author          string    `json:"author" form:"author"`
	CourseCreatedAt time.Time `json:"created_at" form:"created_at"`
}
