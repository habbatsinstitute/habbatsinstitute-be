package dtos

import "time"

type ResEbook struct {
	ID int `json:"id" form:"id"`
	UserID int `json:"user_id" form:"user_id"`
	Ebook string `json:"ebook" form:"ebook"`
	ThumbnailBook string `json:"thumbnail_book" form:"thumbnail_book"`
	Genre string `json:"genre" form:"genre"`
	Title string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	Author string `json:"author" form:"author"`
	BookCreated time.Time `json:"book_created" form:"book_created"`
}

type ResGenre struct {
	ID int `json:"id" form:"id"`
	Genre string `json:"genre" form:"genre"`
}