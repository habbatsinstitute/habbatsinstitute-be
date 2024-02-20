package usecase

import (
	"institute/features/item"
	"institute/features/item/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model item.Repository
}

func New(model item.Repository) item.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResItem {
	var items []dtos.ResItem

	itemsEnt := svc.model.Paginate(page, size)

	for _, item := range itemsEnt {
		var data dtos.ResItem

		if err := smapping.FillStruct(&data, smapping.MapFields(item)); err != nil {
			log.Error(err.Error())
		} 
		
		items = append(items, data)
	}

	return items
}

func (svc *service) FindByID(itemID int) *dtos.ResItem {
	res := dtos.ResItem{}
	item := svc.model.SelectByID(itemID)

	if item == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(item))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newItem dtos.InputItem) *dtos.ResItem {
	item := item.Item{}
	
	err := smapping.FillStruct(&item, smapping.MapFields(newItem))
	if err != nil {
		log.Error(err)
		return nil
	}

	itemID := svc.model.Insert(item)

	if itemID == -1 {
		return nil
	}

	resItem := dtos.ResItem{}
	errRes := smapping.FillStruct(&resItem, smapping.MapFields(newItem))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resItem
}

func (svc *service) Modify(itemData dtos.InputItem, itemID int) bool {
	newItem := item.Item{}

	err := smapping.FillStruct(&newItem, smapping.MapFields(itemData))
	if err != nil {
		log.Error(err)
		return false
	}

	newItem.ID = itemID
	rowsAffected := svc.model.Update(newItem)

	if rowsAffected <= 0 {
		log.Error("There is No Item Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(itemID int) bool {
	rowsAffected := svc.model.DeleteByID(itemID)

	if rowsAffected <= 0 {
		log.Error("There is No Item Deleted!")
		return false
	}

	return true
}