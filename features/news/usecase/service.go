package usecase

import (
	"errors"
	"institute/features/news"
	"institute/features/news/dtos"
	"institute/helpers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model news.Repository
	validator helpers.ValidationInterface
}

func New(model news.Repository, validator helpers.ValidationInterface) news.Usecase {
	return &service {
		model: model,
		validator: validator,
	}
}

func (svc *service) FindAll(page, size int) ([]dtos.ResNews, int64) {
	var newss []dtos.ResNews

	newssEnt := svc.model.Paginate(page, size)
	

	for _, news := range newssEnt {
		var data dtos.ResNews

		if err := smapping.FillStruct(&data, smapping.MapFields(news)); err != nil {
			log.Error(err.Error())
		} 
		
		newss = append(newss, data)
	}
	var totalData int64 = 0
	totalData = svc.model.GetTotalDataNews()

	return newss, totalData
}

func (svc *service) GetTopNews(size int) ([]dtos.ResNews) {
	var newss []dtos.ResNews
	newsEnt := svc.model.GetTopNews(size)

	for _, news := range newsEnt {
		var data dtos.ResNews

		if err := smapping.FillStruct(&data, smapping.MapFields(news)); err != nil {
			log.Error(err.Error())
		} 
		
		newss = append(newss,data)
	}

	return newss
}

func (svc *service) IncrementViews(courseID int) error {
	courseData := svc.model.SelectByID(courseID)

	if courseData == nil {
		return errors.New("data not found")
	}

	courseData.Views++

	if rowsAffected := svc.model.Update(*courseData); rowsAffected <= 0 {
		return errors.New("failed to update views")
	}
	return nil
}

func (svc *service) FindByID(newsID int) *dtos.ResNews {
	res := dtos.ResNews{}

	if err := svc.IncrementViews(newsID); err != nil {
		log.Error(err)
	}
	news := svc.model.SelectByID(newsID)

	if news == nil {
		return nil
	}

	res.ID = news.ID
	res.Title = news.Title
	res.Category = news.Category
	res.Description = news.Description
	res.NewsCreated = news.NewsCreated
	res.Images = news.Images
	res.UserID = news.UserID
	res.Views = news.Views

	return &res
}

func (svc *service) Create(newNews dtos.InputNews,UserID int, file *multipart.FileHeader) (*dtos.ResNews,[]string, error) {
	news := news.News{}

	if errorList, err := svc.ValidateInput(newNews, file); err != nil || len(errorList) > 0 {
		return nil, errorList, err
	}
	url, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil,nil, errors.New("upload image failed")
	}

	news.ID = helpers.NewGenerator().GenerateRandomID()
	news.UserID = UserID
	news.Images = url
	news.Category = newNews.Category
	news.Title = newNews.Title
	news.Description = newNews.Description
	news.NewsCreated = svc.model.GetTimeNow()

	result, err := svc.model.Insert(&news)
	if err != nil {
		log.Error(err)
		return nil,nil, errors.New("failed to create news")
	}

	resNews := dtos.ResNews{}
	resNews.ID = result.ID
	resNews.UserID = result.UserID
	resNews.Category = result.Category
	resNews.Description = result.Description
	resNews.Title = result.Title
	resNews.Images = result.Images

	return &resNews, nil,nil


}

func (svc *service) Modify(newsData dtos.InputNews, newsID int, file *multipart.FileHeader) bool {
	var url string

	if file != nil {
		var err error
		url, err = svc.model.UploadFile(file, newsData.Title)
		if err != nil {
			log.Error("failed upload images")
			return false
		}
	}

	newNews := news.News{
		ID: newsID,
		Category: newsData.Category,
		Title: newsData.Title,
		Description: newsData.Description,
	}

	if file != nil {
		newNews.Images = url
	}

	

	rowsAffected := svc.model.Update(newNews)

	if rowsAffected <= 0 {
		log.Error("There is No News Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(newsID int) bool {
	rowsAffected := svc.model.DeleteByID(newsID)

	if rowsAffected <= 0 {
		log.Error("There is No News Deleted!")
		return false
	}

	return true
}

func (svc *service) FindAllCategory() ([]dtos.ResCategory, error) {
	category, err := svc.model.SelectAllCategory()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (svc *service) SearchNews(title string) ([]dtos.ResNews, error) {
	newsList := svc.model.SearchNewsByTitle(title)

	// convert newslist to dtos
	resNewsList := make([]dtos.ResNews, len(newsList))
	for i, n := range newsList {
		resNewsList[i] = dtos.ResNews{
			ID:    n.ID,
			UserID: n.UserID,
			Title: n.Title,
			Description: n.Description,
			Category: n.Category,
			Images: n.Images,
			NewsCreated: n.NewsCreated,

		}
	}

	return resNewsList, nil
}

func (svc *service) ValidateInput(input dtos.InputNews, fileHeader *multipart.FileHeader) ([]string, error) {
	const (
		minTitleLength      = 19
		maxDescriptionLength = 1999
		maxFileSize          = 2 * 1024 * 1024
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
		isImage := contentType[:5] == "image"

		if !isImage {
			errorList = append(errorList, "file must be image (png, jpg, or jpeg)")
		}

		fileSize, err := io.CopyN(ioutil.Discard, file, maxFileSize+1)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if fileSize > maxFileSize {
			errorList = append(errorList, "file size exceeds the allowed limit (2MB)")
		}
	}

	return errorList, nil
}