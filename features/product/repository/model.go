package repository

import (
	"institute/features/product"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) product.Repository {
	return &model {
		db: db,
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

func (mdl *model) Insert(newProduct product.Product) int64 {
	result := mdl.db.Create(&newProduct)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newProduct.ID)
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
	result := mdl.db.Save(&product)

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