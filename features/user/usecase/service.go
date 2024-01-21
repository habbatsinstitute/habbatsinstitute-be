package usecase

import (
	"institute/features/user"
	"institute/features/user/dtos"
	"institute/helpers"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
)

type service struct {
	model user.Repository
	jwt helpers.JWTInterface
	hash helpers.HashInterface
}

func New(model user.Repository, jwt helpers.JWTInterface, hash helpers.HashInterface) user.Usecase {
	return &service {
		model: model,
		jwt: jwt,
		hash: hash,
	}
}

func (svc *service) FindAll(page, size int) []dtos.ResGetAllUsers {
	var users []dtos.ResGetAllUsers

	usersEnt := svc.model.Paginate(page, size)

	for _, user := range usersEnt {
		var data dtos.ResGetAllUsers

		if err := smapping.FillStruct(&data, smapping.MapFields(user)); err != nil {
			log.Error(err.Error())
		} 
		
		users = append(users, data)
	}

	return users
}

func (svc *service) FindByID(userID int) *dtos.ResUser {
	res := dtos.ResUser{}
	user := svc.model.SelectByID(userID)

	if user == nil {
		return nil
	}

	err := smapping.FillStruct(&res, smapping.MapFields(user))
	if err != nil {
		log.Error(err)
		return nil
	}

	return &res
}

func (svc *service) Modify(userData dtos.InputUser, userID int) bool {
	newUser := user.User{}

	err := smapping.FillStruct(&newUser, smapping.MapFields(userData))
	if err != nil {
		log.Error(err)
		return false
	}

	newUser.ID = userID
	rowsAffected := svc.model.Update(newUser)

	if rowsAffected <= 0 {
		log.Error("There is No User Updated!")
		return false
	}
	
	return true
}

func (svc *service) ModifyUser(userData dtos.UpdateUser, UserID int) bool {
	newData := user.User{}
	
	err := smapping.FillStruct(&newData, smapping.MapFields(userData))
	if err != nil {
		log.Error(err)
		return false
	}

	newData.ID = UserID
	newData.Password = svc.hash.HashPassword(newData.Password)
	rowsAffected := svc.model.Update(newData)

	if rowsAffected <= 0 {
		log.Error("there is no user updated!")
		return false
	}
	return true
}

func (svc *service) Remove(userID int) bool {
	rowsAffected := svc.model.DeleteByID(userID)

	if rowsAffected <= 0 {
		log.Error("There is No User Deleted!")
		return false
	}

	return true
}