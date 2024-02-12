package dtos

type InputNews struct {
	Category 	string	`json:"category" form:"category"`
	Images		string	`json:"images" form:"images"`
	Title		string	`json:"title" form:"title"`
	Description	string	`json:"description" form:"description"`
}

type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}