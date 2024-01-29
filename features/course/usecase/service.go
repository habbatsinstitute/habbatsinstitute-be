package usecase

import (
	"errors"
	"institute/features/course"
	"institute/features/course/dtos"
	"institute/helpers"
	"mime/multipart"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model course.Repository
	jwt helpers.JWTInterface
}

func New(model course.Repository, jwt helpers.JWTInterface) course.Usecase {
	return &service {
		model: model,
		jwt: jwt,
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

func (svc *service) FindByID(courseID int) *dtos.ResCourse {
	res := dtos.ResCourse{}
	course := svc.model.SelectByID(courseID)

	if course == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(course))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newCourse dtos.InputCourse,UserID int, file *multipart.FileHeader) (*dtos.ResCourse, error) {
	course := course.Course{}

	url, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil, errors.New("upload image failed")
	}

	course.ID = helpers.NewGenerator().GenerateRandomID()
	course.UserID = UserID
	course.MediaFile = url
	course.Author = newCourse.Author
	course.Title = newCourse.Title
	course.Description = newCourse.Description

	result, err := svc.model.Insert(&course)
	if err != nil {
		log.Error(err)
		return nil, errors.New("failed to create course")
	}

	resCourse := dtos.ResCourse{}
	resCourse.ID = result.ID
	resCourse.UserID = result.UserID
	resCourse.Author = result.Author
	resCourse.MediaFile = result.MediaFile
	resCourse.Title = result.Title
	resCourse.Description = result.Description

	return &resCourse, nil
}

func (svc *service) Modify(courseData dtos.InputCourse, courseID int) bool {
	newCourse := course.Course{}

	err := smapping.FillStruct(&newCourse, smapping.MapFields(courseData))
	if err != nil {
		log.Error(err)
		return false
	}

	newCourse.ID = courseID
	rowsAffected := svc.model.Update(newCourse)

	if rowsAffected <= 0 {
		log.Error("There is No Course Updated!")
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