package news

import (
	"institute/features/news/dtos"
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []News
	Insert(newNews *News) (*News, error)
	SelectByID(newsID int) *News
	Update(news News) int64
	DeleteByID(newsID int) int64
	UploadFile(fileHeader *multipart.FileHeader, name string) (string, error)
	SelectAllCategory() ([]dtos.ResCategory, error)
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResNews
	FindByID(newsID int) *dtos.ResNews
	Create(newNews dtos.InputNews,UserID int, file *multipart.FileHeader) (*dtos.ResNews, error)
	Modify(newsData dtos.InputNews, newsID int) bool
	Remove(newsID int) bool
	FindAllCategory() ([]dtos.ResCategory, error)
}

type Handler interface {
	GetNewss() echo.HandlerFunc
	NewsDetails() echo.HandlerFunc
	CreateNews() echo.HandlerFunc
	UpdateNews() echo.HandlerFunc
	DeleteNews() echo.HandlerFunc
	GetCategory() echo.HandlerFunc
}
