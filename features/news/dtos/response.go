package dtos

import "time"

type ResNews struct {
	ID			int		`json:"id" form:"id"`
	Category 	string	`json:"category" form:"category"`
	Images		string	`json:"images" form:"images"`
	Title		string	`json:"title" form:"title"`
	Description	string	`json:"description" form:"description"`
	CreatedAt	time.Time `json:"created_at" form:"created_at"`
}

type ResCategory struct {
	ID 		int `json:"id" form:"id"`
	Name 	string `json:"name" form:"name"`
}