package repository

import (
	"institute/features/auth"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.Repository{
	return &model{
		db : db,
	}
}


func (mdl *model) Register(newUser *auth.User) (*auth.User, error){
	result := mdl.db.Table("users").Create(newUser)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}
	return newUser, nil
}

func (mdl *model) Login(username string) (*auth.User, error){
	var user auth.User
	result := mdl.db.Table("users").Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, result.Error
		}
		log.Error(result.Error)
		return nil, result.Error
	}
	return &user, nil
}

func (mdl *model) SelectByUsername(username string) (*auth.User, error) {
	var user auth.User

	result := mdl.db.Table("users").Where("username = ?", username).First(&user)

	if result.Error != nil {
		log.Error(result.Error)
		return nil, result.Error
	}

	return &user, nil
}