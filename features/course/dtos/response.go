package dtos

type ResCourse struct {
	ID 				int 			`json:"id" form:"id"`
	MediaFile			string			`json:"media_file" form:"media_file"`
	Title			string			`json:"title" form:"title"`
	Description		string			`json:"description" form:"description"`
	Author			string			`json:"author" form:"author"`
}
