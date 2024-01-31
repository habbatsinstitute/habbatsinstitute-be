package usecase

import (
	"errors"
	"institute/features/news"
	"institute/features/news/dtos"
	"institute/helpers"
	"mime/multipart"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model news.Repository
}

func New(model news.Repository) news.Usecase {
	return &service {
		model: model,
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

func (svc *service) FindByID(newsID int) *dtos.ResNews {
	res := dtos.ResNews{}
	news := svc.model.SelectByID(newsID)

	if news == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(news))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newNews dtos.InputNews,UserID int, file *multipart.FileHeader) (*dtos.ResNews, error) {
	news := news.News{}
	
	url, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil, errors.New("upload image failed")
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
		return nil, errors.New("failed to create news")
	}

	resNews := dtos.ResNews{}
	resNews.ID = result.ID
	resNews.UserID = result.UserID
	resNews.Category = result.Category
	resNews.Description = result.Description
	resNews.Title = result.Title
	resNews.Images = result.Images

	return &resNews, nil


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