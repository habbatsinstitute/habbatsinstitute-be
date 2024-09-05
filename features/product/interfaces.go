package product

import (
	"institute/features/product/dtos"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Product
	Insert(newProduct *Product) (*Product, error)
	SelectByID(productID int) *Product
	Update(product Product) int64
	DeleteByID(productID int) int64
	UploadFile(fileHeader *multipart.FileHeader, name string) (string, error)
	GetTimeNow() time.Time
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResProduct
	FindByID(productID int) *dtos.ResProduct
	Create(newProduct dtos.InputProduct, UserID int, file *multipart.FileHeader) (*dtos.ResProduct,[]string, error)
	Modify(productData dtos.InputProduct, productID int) bool
	Remove(productID int) bool
	ValidateInput(input dtos.InputProduct, fileHeader *multipart.FileHeader) ([]string, error)
}

type Handler interface {
	GetProducts() echo.HandlerFunc
	ProductDetails() echo.HandlerFunc
	CreateProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
}
