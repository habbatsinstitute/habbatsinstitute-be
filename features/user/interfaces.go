package user

import (
	"institute/features/user/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []User
	SelectByID(userID int) *User
	Update(user User) int64
	DeleteByID(userID int) int64
	GetTotalDataUsers() int64
	FindUsername(username string, page, size int) []User
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResGetAllUsers, int64)
	FindByID(userID int) *dtos.ResUser
	Remove(userID int) bool
	ModifyUser(userData dtos.UpdateUser, UserID int) bool
	MyProfile(UserID int) *dtos.ResMyProfile
	SearchUsersByUsername(username string, page, size int) ([]dtos.ResGetAllUsers, int64)
}

type Handler interface {
	GetUsers() echo.HandlerFunc
	UserDetails() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	UpdateExpiryAccount() echo.HandlerFunc
	MyProfile() echo.HandlerFunc
	SearchNewsByUsername() echo.HandlerFunc
}
