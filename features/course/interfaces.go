package course

import (
	"institute/features/course/dtos"
	"mime/multipart"
	"time"

	"github.com/labstack/echo/v4"
)

type Repository interface {
	Paginate(page, size int, search dtos.Search) []Course
	Insert(newCourse *Course) (*Course, error)
	SelectByID(courseID int) *Course
	Update(course Course) int64
	DeleteByID(courseID int) int64
	UploadFile(fileHeader *multipart.FileHeader, name string) (string, error)
	GetTotalDataCourseBySearchAndFilter(search dtos.Search) int64
	GetTotalDataCourse() int64
	GetTimeNow() time.Time 
}

type Usecase interface {
	FindAll(page, size int, search dtos.Search) ([]dtos.ResCourse, int64)
	FindByID(courseID int) *dtos.ResCourse
	Create(newCourse dtos.InputCourse,UserID int, file *multipart.FileHeader) (*dtos.ResCourse,[]string, error)
	Modify(courseData dtos.InputCourse, courseID int, file *multipart.FileHeader) bool
	Remove(courseID int) bool
	IncrementViews(courseID int) error
}

type Handler interface {
	GetCourses() echo.HandlerFunc
	CourseDetails() echo.HandlerFunc
	CreateCourse() echo.HandlerFunc
	UpdateCourse() echo.HandlerFunc
	DeleteCourse() echo.HandlerFunc
}
