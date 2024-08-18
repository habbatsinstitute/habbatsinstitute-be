package dtos

import "time"

type ResEbook struct {
	ID int `json:"id" form:"id"`
	UserID int `json:"user_id" form:"user_id"`
	Ebook string `json:"ebook" form:"ebook"`
	ThumbnailBook string `json:"thumbnail_book" form:"thumbnail_book"`
	GenreEbook string `json:"genre" form:"genre"`
	TitleEbook string `json:"title" form:"title"`
	DescriptionEbook string `json:"description" form:"description"`
	AuthorEbook string `json:"author" form:"author"`
	BookCreated time.Time `json:"book_created" form:"book_created"`
}

type ResGenre struct {
	ID int `json:"id" form:"id"`
	Genre string `json:"genre" form:"genre"`
}