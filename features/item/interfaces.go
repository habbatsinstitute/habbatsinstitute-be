package item

import (
	"institute/features/item/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Item
	Insert(newItem Item) int64
	SelectByID(itemID int) *Item
	Update(item Item) int64
	DeleteByID(itemID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResItem
	FindByID(itemID int) *dtos.ResItem
	Create(newItem dtos.InputItem) *dtos.ResItem
	Modify(itemData dtos.InputItem, itemID int) bool
	Remove(itemID int) bool
}

type Handler interface {
	GetItems() echo.HandlerFunc
	ItemDetails() echo.HandlerFunc
	CreateItem() echo.HandlerFunc
	UpdateItem() echo.HandlerFunc
	DeleteItem() echo.HandlerFunc
}
