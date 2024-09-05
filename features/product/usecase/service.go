package usecase

import (
	"errors"
	"institute/features/product"
	"institute/features/product/dtos"
	"institute/helpers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model product.Repository
	validator helpers.ValidationInterface
}

func New(model product.Repository, validator helpers.ValidationInterface) product.Usecase {
	return &service {
		model: model,
		validator: validator,
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

func (svc *service) Create(newProduct dtos.InputProduct, UserID int, file *multipart.FileHeader) (*dtos.ResProduct,[]string, error) {
	product := product.Product{}

	if errorList, err := svc.ValidateInput(newProduct, file); err != nil || len(errorList) > 0 {
		return nil, errorList, err
	}
	
	url, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil, nil, errors.New("upload image failed")
	}

	product.ID = helpers.NewGenerator().GenerateRandomID()
	product.UserID = UserID
	product.Name = newProduct.Name
	product.Images = url
	product.Description = newProduct.Description
	product.Composition = newProduct.Composition
	product.PomTR = newProduct.PomTR
	product.MarketPlace = newProduct.MarketPlace
	product.Whatsapp = newProduct.Whatsapp
	product.Price = newProduct.Price
	product.Quantity = newProduct.Quantity
	product.ProductCreated = svc.model.GetTimeNow()

	result, err := svc.model.Insert(&product)
	if err != nil {
		log.Error(err)
		return nil, nil, errors.New("fail to create product")
	}

	resProduct := dtos.ResProduct{}
	resProduct.Composition = result.Composition
	resProduct.MarketPlace = result.MarketPlace
	resProduct.Price = result.Price
	resProduct.Quantity = result.Quantity
	resProduct.Images = result.Images
	resProduct.PomTR = result.PomTR
	resProduct.Name = result.Name
	resProduct.Whatsapp = result.Whatsapp
	resProduct.Description = result.Description

	return &resProduct, nil, nil
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

func (svc *service) ValidateInput(input dtos.InputProduct, fileHeader *multipart.FileHeader) ([]string, error){
	const (
		maxFileSize = 2 * 1024 * 1024
	)

	var errorList []string

	if errMap := svc.validator.ValidateRequest(input); errMap != nil {
		errorList = append(errorList, errMap...)
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