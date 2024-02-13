package news

import (
	"institute/features/news/dtos"
	"mime/multipart"
	"time"

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
	GetTotalDataNews() int64
	SearchNewsByTitle(title string) []News
	GetTimeNow() time.Time
}

type Usecase interface {
	FindAll(page, size int) ([]dtos.ResNews, int64) 
	FindByID(newsID int) *dtos.ResNews
	Create(newNews dtos.InputNews,UserID int, file *multipart.FileHeader) (*dtos.ResNews,[]string, error)
	Modify(newsData dtos.InputNews, newsID int, file *multipart.FileHeader) bool
	Remove(newsID int) bool
	FindAllCategory() ([]dtos.ResCategory, error)
	SearchNews(title string) ([]dtos.ResNews, error)
	IncrementViews(courseID int) error
}

type Handler interface {
	GetNews() echo.HandlerFunc
	NewsDetails() echo.HandlerFunc
	CreateNews() echo.HandlerFunc
	UpdateNews() echo.HandlerFunc
	DeleteNews() echo.HandlerFunc
	GetCategory() echo.HandlerFunc
	SearchNewsByTitle() echo.HandlerFunc
}
