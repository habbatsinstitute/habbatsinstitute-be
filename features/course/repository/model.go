package repository

import (
	"fmt"
	"institute/config"
	"institute/features/course"
	"institute/features/course/dtos"
	"institute/helpers"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
	cdn *cloudinary.Cloudinary
	config *config.ProgramConfig
}

func New(db *gorm.DB, cdn *cloudinary.Cloudinary, config *config.ProgramConfig) course.Repository {
	return &model {
		db: db,
		cdn: cdn,
		config: config,
	}
}

func (mdl *model) Paginate(page, size int, search dtos.Search) []course.Course {
	var courses []course.Course

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Where("title LIKE ?",search.Title+"%").Find(&courses)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return courses
}

func (mdl *model) Insert(newCourse *course.Course) (*course.Course, error) {
	result := mdl.db.Create(&newCourse)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}

	return newCourse, nil
}

func (mdl *model) SelectByID(courseID int) *course.Course {
	var course course.Course
	result := mdl.db.First(&course, courseID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &course
}

func (mdl *model) Update(course course.Course) int64 {
	result := mdl.db.Save(&course)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(courseID int) int64 {
	result := mdl.db.Delete(&course.Course{}, courseID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) UploadFile(fileHeader *multipart.FileHeader, name string) (string, error){
	file := helpers.OpenFileHeader(fileHeader)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cfg := mdl.config.CDN_Folder_Name

	resp, err := mdl.cdn.Upload.Upload(ctx, file, uploader.UploadParams{
		Folder: cfg,
		PublicID: name,
	})

	if err != nil {
		fmt.Println(err.Error())
		return "", nil
	}
	return resp.SecureURL, nil
}

func (mdl *model) GetTotalDataVacanciesBySearchAndFilter(search dtos.Search) int64 {
	var totalData int64

	result := mdl.db.Table("courses").
		Where("title LIKE ?", "%"+search.Title+"%").Count(&totalData)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return totalData
}

func (mdl *model) GetTotalDataCourse() int64 {
	var totalData int64

	result := mdl.db.Table("courses").Where("deleted_at IS NULL").Count(&totalData)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return totalData
}