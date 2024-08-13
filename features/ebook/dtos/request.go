package dtos

type InputEbook struct {
	Ebook string `json:"ebook" form:"ebook"`
	ThumbnailBook string `json:"thumbnail_book" form:"thumbnail_book"`
	Genre string `json"genre" form:"genre"`
	Title string `json"title" form:"title"`
	Description string `json"description" form:"description"`
	Author string `json"author" form:"author"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}