package usecase

import (
	"errors"
	"fmt"
	"institute/features/ebook"
	"institute/features/ebook/dtos"
	"institute/helpers"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model ebook.Repository
	validator helpers.ValidationInterface
}

func New(model ebook.Repository, validator helpers.ValidationInterface) ebook.Usecase {
	return &service {
		model: model,
		validator: validator,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResEbook {
	var ebooks []dtos.ResEbook

	ebooksEnt := svc.model.Paginate(page, size)

	for _, ebook := range ebooksEnt {
		var data dtos.ResEbook

		if err := smapping.FillStruct(&data, smapping.MapFields(ebook)); err != nil {
			log.Error(err.Error())
		} 
		
		ebooks = append(ebooks, data)
	}

	return ebooks
}

func (svc *service) FindByID(ebookID int) *dtos.ResEbook {
	res := dtos.ResEbook{}
	ebook := svc.model.SelectByID(ebookID)

	if ebook == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(ebook))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Create(newEbook dtos.InputEbook,UserID int, file *multipart.FileHeader, thumbnail *multipart.FileHeader) (*dtos.ResEbook, []string, error) {
	ebook := ebook.Ebook{}

	if errorList, err := svc.ValidateInput(newEbook, file); err != nil || len(errorList) > 0 {
		return nil, errorList, err
	}
	urlFile, err := svc.model.UploadFile(file, "")
	if err != nil {
		return nil, nil, errors.New("failed to upload file")
	}

	urlThumbnail, err := svc.model.UploadFile(thumbnail, "")
	if err != nil {
		return nil, nil, errors.New("failed to upload thumbnail")
	}

	ebook.ID = helpers.NewGenerator().GenerateRandomID()
	ebook.UserID = UserID
	ebook.ThumbnailBook = urlThumbnail
	ebook.Ebook = urlFile
	ebook.TitleEbook = newEbook.TitleEbook
	ebook.DescriptionEbook = newEbook.DescriptionEbook
	ebook.GenreEbook = newEbook.GenreEbook
	ebook.AuthorEbook = newEbook.AuthorEbook
	ebook.BookCreated = svc.model.GetTimeNow()

	result, err := svc.model.Insert(&ebook)
	if err != nil {
		log.Error(err)
		return nil, nil, errors.New("failed to create ebook")
	}

	resEboook := dtos.ResEbook{}
	resEboook.ID = result.ID
	resEboook.Ebook = result.Ebook
	resEboook.ThumbnailBook = result.ThumbnailBook
	resEboook.TitleEbook = result.TitleEbook
	resEboook.DescriptionEbook = result.DescriptionEbook
	resEboook.GenreEbook = result.GenreEbook
	resEboook.AuthorEbook = result.AuthorEbook

	fmt.Println("services: insert data: ", &resEboook)

	return &resEboook, nil, nil
}

func (svc *service) Modify(ebookData dtos.InputEbook, ebookID int, file *multipart.FileHeader, thumbnail *multipart.FileHeader) bool {
	var urlFile string
	var urlThumbnail string

	if file != nil {
		var err error
		urlFile, err = svc.model.UploadFile(file, ebookData.Ebook)
		if err != nil {
			log.Error("failed to upload file")
			return false
		}
	}

	if thumbnail != nil {
		var err error
		urlThumbnail, err = svc.model.UploadFile(thumbnail, ebookData.ThumbnailBook)
		if err != nil {
			log.Error("failed to upload thumbnail")
			return false
		}
	}

	newEbook := ebook.Ebook{
		ID: ebookID,
		GenreEbook: ebookData.GenreEbook,
		DescriptionEbook: ebookData.DescriptionEbook,
		AuthorEbook: ebookData.AuthorEbook,
		TitleEbook: ebookData.TitleEbook,
	}
	if file != nil {
		newEbook.Ebook = urlFile
	}
	if thumbnail != nil {
		newEbook.ThumbnailBook = urlThumbnail
	}

	rowsAffected := svc.model.Update(newEbook)

	if rowsAffected <= 0 {
		log.Error("there is no book updated!")
		return false
	}
	return true
}

func (svc *service) Remove(ebookID int) bool {
	rowsAffected := svc.model.DeleteByID(ebookID)

	if rowsAffected <= 0 {
		log.Error("There is No Ebook Deleted!")
		return false
	}

	return true
}

func (svc *service) ValidateInput(input dtos.InputEbook, fileHeader *multipart.FileHeader) ([]string, error){
	const(
		minTitleLength = 19
		maxDescriptionLength = 4999
		maxFileSize = 100 * 1024 * 1024
		maxThumbnailSize = 2 * 1024 * 1024
		maxAuthorLength = 30
	)

	var errorList []string
	
	if errMap := svc.validator.ValidateRequest(input); errMap != nil {
		errorList = append(errorList, errMap...)
	}

	if len(input.TitleEbook ) <= minTitleLength {
		errorList = append(errorList, "Title must be at least 20 characters")
	}
	if len(input.DescriptionEbook) >= maxDescriptionLength {
		errorList = append(errorList, "description maximum length must be at least 5000 characters")
	}
	if len(input.AuthorEbook) >= maxAuthorLength {
		errorList = append(errorList, "author maximum length must be at least 30 characters")
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

		isPDF := contentType == "application/pdf"
		if !isPDF{
			errorList = append(errorList, "file must be a PDF document")
		}

		fileSize, err := io.CopyN(ioutil.Discard, file, maxFileSize+1)
		if err != nil && err != io.EOF {
			return nil, err
		}

		if fileSize > maxFileSize {
			errorList = append(errorList, "file size exceeds the allowed limit (100MB)")
		}
	}
	return errorList, nil
}