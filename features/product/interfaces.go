package product

import (
	"institute/features/product/dtos"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Product
	Insert(newProduct Product) int64
	SelectByID(productID int) *Product
	Update(product Product) int64
	DeleteByID(productID int) int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResProduct
	FindByID(productID int) *dtos.ResProduct
	Create(newProduct dtos.InputProduct) *dtos.ResProduct
	Modify(productData dtos.InputProduct, productID int) bool
	Remove(productID int) bool
}

type Handler interface {
	GetProducts() echo.HandlerFunc
	ProductDetails() echo.HandlerFunc
	CreateProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}
