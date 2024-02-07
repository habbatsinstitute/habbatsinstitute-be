package dtos

import "time"

type ResUser struct {
	Username 	   string		`json:"username" form:"username"`
	ExpiryDate 	   time.Time	`json:"expiry_date" form:"expiry_date"`
}

type ResUpdateUser struct {
	Username 	   string		`json:"username" form:"username"`	
	Password       string		`json:"password" form:"password"`
	ExpiryDate 	   time.Time	`json:"expiry_date" form:"expiry_date"`
}

type ResGetAllUsers struct {
	ID 			   int			`json:"id" form:"id"`
	Username 	   string		`json:"username" form:"username"`
	RoleID		   int		`json:"role_id" form:"role_id"`
	ExpiryDate 	   time.Time	`json:"expiry_date" form:"expiry_date"`
}

type ResMyProfile struct {
	ID int 					`json:"id" form:"id"`
	Username string 		`json:"username" form:"username"`
	RoleID int 				`json:"role_id" form:"role_id"`
	ExpiryDate time.Time 	`json:"expiry_date" form:"expiry_date"`
}