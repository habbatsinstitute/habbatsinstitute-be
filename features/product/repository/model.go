package repository

import (
	"context"
	"fmt"
	"institute/config"
	"institute/features/product"
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

func New(db *gorm.DB, cdn *cloudinary.Cloudinary, config *config.ProgramConfig) product.Repository {
	return &model {
		db: db,
		cdn: cdn,
		config: config,
	}
}

func (mdl *model) Paginate(page, size int) []product.Product {
	var products []product.Product

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&products)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return products
}

func (mdl *model) Insert(newProduct *product.Product) (*product.Product, error) {
	result := mdl.db.Create(&newProduct)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}

	return newProduct, nil
}

func (mdl *model) SelectByID(productID int) *product.Product {
	var product product.Product
	result := mdl.db.First(&product, productID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &product
}

func (mdl *model) Update(product product.Product) int64 {
	result := mdl.db.Updates(&product)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(productID int) int64 {
	result := mdl.db.Delete(&product.Product{}, productID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) UploadFile(fileHeader *multipart.FileHeader, name string) (string, error) {
	file := helpers.OpenFileHeader(fileHeader)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cfg := mdl.config.CDN_FOLDER_PRODUCT

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

func (mdl *model) GetTimeNow() time.Time {
	wibLocation, _ := time.LoadLocation("Asia/Jakarta")

	return time.Now().In(wibLocation)
}