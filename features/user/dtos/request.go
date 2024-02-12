package dtos

import "time"

type InputUser struct {
	Username 	   string		`json:"username" form:"username"`	
}

type UpdateUser struct {
	Username 	   string		`json:"username" form:"username"`	
	Password       string		`json:"password" form:"password"`
	ExpiryDate 	   time.Time	`json:"expiry_date" form:"expiry_date"`
}


type Pagination struct {
	Page int `query:"page"`
	Size int `query:"page_size"`
}