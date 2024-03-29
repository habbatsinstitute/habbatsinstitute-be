package repository

import (
	"context"
	"fmt"
	"institute/config"
	"institute/features/news"
	"institute/features/news/dtos"
	"institute/helpers"
	"mime/multipart"
	"time"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
	cdn *cloudinary.Cloudinary
	config *config.ProgramConfig
}

func New(db *gorm.DB, cdn *cloudinary.Cloudinary, config *config.ProgramConfig) news.Repository {
	return &model {
		db: db,
		cdn: cdn,
		config: config,
	}
}

func (mdl *model) Paginate(page, size int) []news.News {
	var news []news.News

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&news)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return news
}

func (mdl *model) GetTopNews(size int) []news.News {
	var news []news.News

	result := mdl.db.Order("views DESC").Limit(size).Find(&news)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return news
}

func (mdl *model) Insert(newNews *news.News) (*news.News, error) {
	result := mdl.db.Create(&newNews)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}

	return newNews, nil
}

func (mdl *model) SelectByID(newsID int) *news.News {
	var news news.News
	result := mdl.db.First(&news, newsID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &news
}

func (mdl *model) SelectAllCategory() ([]dtos.ResCategory, error) {
	var Category []dtos.ResCategory

	if err := mdl.db.Table("categories").Find(&Category).Error; err != nil {
		return nil, err
	}

	return Category, nil
}

func (mdl *model) Update(news news.News) int64 {
	result := mdl.db.Updates(&news)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(newsID int) int64 {
	result := mdl.db.Delete(&news.News{}, newsID)
	
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
	cfg := mdl.config.CDN_FOLDER_ARTICLES

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

func (mdl *model) GetTotalDataNews() int64 {
	var totalData int64

	result := mdl.db.Table("news").Where("deleted_at IS NULL").Count(&totalData)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return totalData
}

func (mdl *model) SearchNewsByTitle(title string) []news.News{
	var news []news.News

	result := mdl.db.Where("title LIKE ?", "%"+title+"%").Find(&news)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	return news
}

func (mdl *model) GetTimeNow() time.Time {
	wibLocation, _ := time.LoadLocation("Asia/Jakarta")

	return time.Now().In(wibLocation)
}