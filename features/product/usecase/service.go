package usecase

import (
	"institute/features/product"
	"institute/features/product/dtos"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model product.Repository
}

func New(model product.Repository) product.Usecase {
	return &service {
		model: model,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResProduct {
	var products []dtos.ResProduct

	productsEnt := svc.model.Paginate(page, size)

	for _, product := range productsEnt {
		var data dtos.ResProduct

		if err := smapping.FillStruct(&data, smapping.MapFields(product)); err != nil {
			log.Error(err.Error())
		} 
		
		products = append(products, data)
	}

	return products
}

func (svc *service) FindByID(productID int) *dtos.ResProduct {
	res := dtos.ResProduct{}
	product := svc.model.SelectByID(productID)

	if product == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(product))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newProduct dtos.InputProduct) *dtos.ResProduct {
	product := product.Product{}
	
	err := smapping.FillStruct(&product, smapping.MapFields(newProduct))
	if err != nil {
		log.Error(err)
		return nil
	}

	productID := svc.model.Insert(product)

	if productID == -1 {
		return nil
	}

	resProduct := dtos.ResProduct{}
	errRes := smapping.FillStruct(&resProduct, smapping.MapFields(newProduct))
	if errRes != nil {
		log.Error(errRes)
		return nil
	}

	return &resProduct
}

func (svc *service) Modify(productData dtos.InputProduct, productID int) bool {
	newProduct := product.Product{}

	err := smapping.FillStruct(&newProduct, smapping.MapFields(productData))
	if err != nil {
		log.Error(err)
		return false
	}

	newProduct.ID = productID
	rowsAffected := svc.model.Update(newProduct)

	if rowsAffected <= 0 {
		log.Error("There is No Product Updated!")
		return false
	}
	
	return true
}

func (svc *service) Remove(productID int) bool {
	rowsAffected := svc.model.DeleteByID(productID)

	if rowsAffected <= 0 {
		log.Error("There is No Product Deleted!")
		return false
	}

	return true
}