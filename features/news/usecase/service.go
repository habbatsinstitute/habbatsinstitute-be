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

func (svc *service) FindAll(page, size int) []dtos.ResNews {
	var newss []dtos.ResNews

	newssEnt := svc.model.Paginate(page, size)

	for _, news := range newssEnt {
		var data dtos.ResNews

		if err := smapping.FillStruct(&data, smapping.MapFields(news)); err != nil {
			log.Error(err.Error())
		} 
		
		newss = append(newss, data)
	}

	return newss
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

func (svc *service) Create(newNews dtos.InputNews, file *multipart.FileHeader) (*dtos.ResNews, error) {
	news := news.News{}
	
	url, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil, errors.New("upload image failed")
	}

	news.ID = helpers.NewGenerator().GenerateRandomID()
	news.Images = url
	news.Category = newNews.Category
	news.Title = newNews.Title
	news.Description = newNews.Description

	result, err := svc.model.Insert(&news)
	if err != nil {
		log.Error(err)
		return nil, errors.New("failed to create news")
	}

	resNews := dtos.ResNews{}
	resNews.ID = result.ID
	resNews.Category = result.Category
	resNews.Description = result.Description
	resNews.Title = result.Title
	resNews.Images = result.Images

	return &resNews, nil


}

func (svc *service) Modify(newsData dtos.InputNews, newsID int) bool {
	newNews := news.News{}

	err := smapping.FillStruct(&newNews, smapping.MapFields(newsData))
	if err != nil {
		log.Error(err)
		return false
	}

	newNews.ID = newsID
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