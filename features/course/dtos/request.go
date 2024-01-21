package dtos

type InputCourse struct {
	MediaFile			string			`json:"media_file" form:"media_file"`
	Title			string			`json:"title" form:"title"`
	Description		string			`json:"description" form:"description"`
	Author			string			`json:"author" form:"author"`
}

type Pagination struct {
	Page int `query:"page"`
	PageSize int `query:"page_size"`
}

type Search struct {
	Title string `query:"title"`
}