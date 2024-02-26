package repository

import (
	"institute/features/user"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.Repository {
	return &model {
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) []user.User {
	var users []user.User

	offset := (page - 1) * size

	result := mdl.db.Offset(offset).Limit(size).Where("role_id = 1").Find(&users)
	
	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return users
}


func (mdl *model) SelectByID(userID int) *user.User {
	var user user.User
	result := mdl.db.First(&user, userID)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}

	return &user
}

func (mdl *model) Update(user user.User) int64 {
	result := mdl.db.Updates(&user)

	if result.Error != nil {
		log.Error(result.Error)
	}

	return result.RowsAffected
}

func (mdl *model) DeleteByID(userID int) int64 {
	result := mdl.db.Delete(&user.User{}, userID)
	
	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}

	return result.RowsAffected
}

func (mdl *model) GetTotalDataUsers() int64 {
	var totalData int64

	result := mdl.db.Table("users").Where("deleted_at IS NULL").Count(&totalData)

	if result.Error != nil {
		log.Error(result.Error)
		return 0
	}
	return totalData
}

func (mdl *model) FindUsername(username string) []user.User{
	var users []user.User

	result := mdl.db.Where("username LIKE ?", "%"+username+"%").Find(&users)

	if result.Error != nil {
		log.Error(result.Error)
		return nil
	}
	return users
}
