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
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResGetAllUsers, int64)
	FindByID(userID int) *dtos.ResUser
	Modify(userData dtos.InputUser, userID int) bool
	Remove(userID int) bool
	ModifyUser(userData dtos.UpdateUser, UserID int) bool
	MyProfile(UserID int) *dtos.ResMyProfile
}

type Handler interface {
	GetUsers() echo.HandlerFunc
	UserDetails() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	UpdateExpiryAccount() echo.HandlerFunc
	MyProfile() echo.HandlerFunc
}
