package ebook

import (
	"institute/features/ebook/dtos"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int) []Ebook
	Insert(newEbook *Ebook) (*Ebook, error)
	SelectByID(ebookID int) *Ebook
	Update(ebook Ebook) int64
	DeleteByID(ebookID int) int64
	SelectAllGenre(genreName string) ([]dtos.ResGenre, error)
	SearchBookByTitle(title string) []Ebook
	GetTimeNow() time.Time
	UploadFile(fileHeader *multipart.FileHeader, name string) (string, error)
	GetTotalDataEbook() int64
}

type Usecase interface {
	FindAll(page, size int) []dtos.ResEbook
	FindByID(ebookID int) *dtos.ResEbook
	Create(newEbook dtos.InputEbook,UserID int, file *multipart.FileHeader, thumbnail *multipart.FileHeader) (*dtos.ResEbook, []string, error)
	Modify(ebookData dtos.InputEbook, ebookID int, file *multipart.FileHeader, thumbnail *multipart.FileHeader) bool
	Remove(ebookID int) bool
	ValidateInput(input dtos.InputEbook, fileHeader *multipart.FileHeader) ([]string, error)
}

type Handler interface {
	GetEbooks() echo.HandlerFunc
	EbookDetails() echo.HandlerFunc
	CreateEbook() echo.HandlerFunc
	// UpdateEbook() echo.HandlerFunc
	DeleteEbook() echo.HandlerFunc
}
