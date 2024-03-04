package usecase

import (
	"errors"
	"institute/features/course"
	"institute/features/course/dtos"
	"institute/helpers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model course.Repository
	jwt helpers.JWTInterface
	validator helpers.ValidationInterface
}

func New(model course.Repository, jwt helpers.JWTInterface, validator helpers.ValidationInterface) course.Usecase {
	return &service {
		model: model,
		jwt: jwt,
		validator: validator,
	}
}

func (svc *service) FindAll(page, size int, search dtos.Search) ([]dtos.ResCourse, int64) {
	var courses []dtos.ResCourse

	coursesEnt := svc.model.Paginate(page, size, search)

	for _, course := range coursesEnt {
		var data dtos.ResCourse

		if err := smapping.FillStruct(&data, smapping.MapFields(course)); err != nil {
			log.Error(err.Error())
		} 
		
		courses = append(courses, data)
	}
	var totalData int64 = 0
	if search.Title != ""{
		totalData = svc.model.GetTotalDataCourseBySearchAndFilter(search)
	}else {
		totalData = svc.model.GetTotalDataCourse()
	}

	return courses, totalData
}

func (svc *service) IncrementViews(courseID int) error {
	courseData := svc.model.SelectByID(courseID)
	if courseData == nil {
		return errors.New("course not found")
	}
	courseData.Views++

	if rowsAffected := svc.model.Update(*courseData); rowsAffected <= 0 {
		return errors.New("falied to update views")
	}
	return nil
}

func (svc *service) FindByID(courseID int) *dtos.ResCourse {
	res := dtos.ResCourse{}

	if err := svc.IncrementViews(courseID); err != nil {
		log.Error(err)
	}
	courseData := svc.model.SelectByID(courseID)

	if courseData == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(courseData))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newCourse dtos.InputCourse,UserID int, file *multipart.FileHeader) (*dtos.ResCourse,[]string, error) {
	course := course.Course{}

	if errorList, err := svc.ValidateInput(newCourse, file); err != nil || len(errorList) > 0 {
		return nil, errorList, err
	}

	url, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil, nil, errors.New("upload image failed")
	}

	course.ID = helpers.NewGenerator().GenerateRandomID()
	course.UserID = UserID
	course.MediaFile = url
	course.Author = newCourse.Author
	course.Title = newCourse.Title
	course.Description = newCourse.Description
	course.CourseCreatedAt = svc.model.GetTimeNow()

	result, err := svc.model.Insert(&course)
	if err != nil {
		log.Error(err)
		return nil, nil, errors.New("failed to create course")
	}

	resCourse := dtos.ResCourse{}
	resCourse.ID = result.ID
	resCourse.UserID = result.UserID
	resCourse.Author = result.Author
	resCourse.MediaFile = result.MediaFile
	resCourse.Title = result.Title
	resCourse.Description = result.Description
	resCourse.CourseCreatedAt = result.CourseCreatedAt

	return &resCourse, nil, nil
}

func (svc *service) Modify(courseData dtos.InputCourse, courseID int, file *multipart.FileHeader) bool {
	var url string

	if file != nil {
		var err error
		url, err = svc.model.UploadFile(file, courseData.Title)
		if err != nil {
			log.Error("failed upload images")
			return false
		}
	}
	newCourse := course.Course{
		ID: courseID,
		Title: courseData.Title,
		Description: courseData.Description,
		Author: courseData.Author,
	}

	if file != nil {
		newCourse.MediaFile = url
	}

	rowsAffected := svc.model.Update(newCourse)

	if rowsAffected < 0 {
		log.Error("there is no course updated!")
		return false
	}
	return true
}

func (svc *service) Remove(courseID int) bool {
	rowsAffected := svc.model.DeleteByID(courseID)

	if rowsAffected <= 0 {
		log.Error("There is No Course Deleted!")
		return false
	}

	return true
}

func (svc *service) ValidateInput(input dtos.InputCourse, fileHeader *multipart.FileHeader) ([]string, error) {
	const (
		minTitleLength      = 19
		maxDescriptionLength = 1999
		maxAuthorLength      = 30
		maxFileSize          = 300 * 1024 * 1024
	)

	var errorList []string

	if errMap := svc.validator.ValidateRequest(input); errMap != nil {
		errorList = append(errorList, errMap...)
	}

	if len(input.Title) <= minTitleLength {
		errorList = append(errorList, "title must be at least 20 characters")
	}

	if len(input.Description) >= maxDescriptionLength {
		errorList = append(errorList, "description must be lower than 2000 characters")
	}

	if len(input.Author) >= maxAuthorLength {
		errorList = append(errorList, "author must be lower than 30 characters")
	}

	if fileHeader != nil {
		file, err := fileHeader.Open()
		if err != nil {
			return nil, err
		}
		defer file.Close()

		buffer := make([]byte, 512)
		_, err = file.Read(buffer)

		if err != nil {
			return nil, err
		}

		contentType := http.DetectContentType(buffer)
		isVideo := isVideoContentType(contentType)

		if !isVideo {
			errorList = append(errorList, "file must be a video (mp4)")
		}

		fileSize, err := io.CopyN(ioutil.Discard, file, maxFileSize+1)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if fileSize > maxFileSize {
			errorList = append(errorList, "file size exceeds the allowed limit (300MB)")
		}
	}

	return errorList, nil
}

func isVideoContentType(contentType string) bool {
	// Add more video formats if needed
	switch contentType {
	case "video/mp4":
		return true
	default:
		return false
	}
}

