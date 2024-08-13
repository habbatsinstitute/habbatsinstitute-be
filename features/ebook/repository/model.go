package repository

import (
	"context"
	"fmt"
	"institute/config"
	"institute/features/ebook"
	"institute/features/ebook/dtos"
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

func New(db *gorm.DB, cdn *cloudinary.Cloudinary, config *config.ProgramConfig) ebook.Repository {
	return &model {
		db: db,
		cdn: cdn,
		config: config,
	}
}

func (mdl *model) Paginate(page, size int) []ebook.Ebook {
	var ebooks []ebook.Ebook

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&ebooks)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return ebooks
}

func (mdl *model) Insert(newEbook *ebook.Ebook) (*ebook.Ebook, error) {
	result := mdl.db.Create(&newEbook)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}

	return newEbook, nil
}

func (mdl *model) SelectByID(ebookID int) *ebook.Ebook {
	var ebook ebook.Ebook
	result := mdl.db.First(&ebook, ebookID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &ebook
}

func (mdl *model) SelectAllGenre(genreName string) ([]dtos.ResGenre, error){
	var genre []dtos.ResGenre
	result := mdl.db.First(&genreName, genre)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	} 
	return genre, nil
}

func (mdl *model) Update(ebook ebook.Ebook) int64 {
	result := mdl.db.Updates(&ebook)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(ebookID int) int64 {
	result := mdl.db.Delete(&ebook.Ebook{}, ebookID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) SearchBookByTitle(title string) []ebook.Ebook{
	var ebook []ebook.Ebook

	result := mdl.db.Where("title LIKE ?", "%"+title+"%").Find(&ebook)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return ebook
}

func (mdl *model) GetTimeNow() time.Time {
	wibLocation, _ := time.LoadLocation("Asia/Jakarta")

	return time.Now().In(wibLocation)
}

func (mdl *model) UploadFile(fileHeader *multipart.FileHeader, name string) (string, error){
	file := helpers.OpenFileHeader(fileHeader)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg := mdl.config.CDN_FOLDER_EBOOKS

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

func (mdl *model) GetTotalDataEbook() int64 {
	var totalData int64

	result := mdl.db.Table("ebooks").Where("deleted_at IS NULL").Count(&totalData)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}
	return totalData
}