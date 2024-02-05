package dtos

import "time"

type RequestLogin struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

type InputUser struct {
	Username   string    `json:"username" form:"username" validate:"required"`
	Password   string    `json:"password" form:"password" validate:"required,alphanum,min=8"`
	ExpiryDate time.Time `json:"expiry_date" form:"expiry_date"`
}

type RefreshJWT struct {
	RefreshToken string `json:"refresh_token" form:"refresh_token" validate:"required"`
}