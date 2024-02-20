package repository

import (
	"institute/features/item"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) item.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []item.Item {
	var items []item.Item

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Find(&items)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return items
}

func (mdl *model) Insert(newItem item.Item) int64 {
	result := mdl.db.Create(&newItem)

	if result.Error != nil {
		log.Error(result.Error)
		return -1
	}

	return int64(newItem.ID)
}

func (mdl *model) SelectByID(itemID int) *item.Item {
	var item item.Item
	result := mdl.db.First(&item, itemID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &item
}

func (mdl *model) Update(item item.Item) int64 {
	result := mdl.db.Save(&item)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(itemID int) int64 {
	result := mdl.db.Delete(&item.Item{}, itemID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}