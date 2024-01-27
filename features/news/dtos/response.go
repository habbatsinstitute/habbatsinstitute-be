package dtos

type ResNews struct {
	ID			int		`json:"id" form:"id"`
	Category 	string	`json:"category" form:"category"`
	Images		string	`json:"images" form:"images"`
	Title		string	`json:"title" form:"title"`
	Description	string	`json:"description" form:"description"`
}

