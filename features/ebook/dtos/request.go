package dtos

type InputEbook struct {
	Ebook string `json:"ebook" form:"ebook"`
	ThumbnailBook string `json:"thumbnail_book" form:"thumbnail_book"`
	GenreEbook string `json:"genre" form:"genre"`
	TitleEbook string `json:"title" form:"title"`
	DescriptionEbook string `json:"description" form:"description"`
	AuthorEbook string `json:"author" form:"author"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}