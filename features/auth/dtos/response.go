package dtos

import "time"

type ResUser struct {
	ID         int       `json:"id"`
	RoleID     int       `json:"role_id"`
	Username   string    `json:"username"`
	ExpiryDate time.Time `json:"expiry_date"`
}

type LoginResponse struct {
	Username     string `json:"username"`
	RoleID       int    `json:"role_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ResJWT struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}