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

func (svc *service) FindAll(page, size int) ([]dtos.ResGetAllUsers, int64) {
	var users []dtos.ResGetAllUsers

	usersEnt := svc.model.Paginate(page, size)

	for _, user := range usersEnt {
		var data dtos.ResGetAllUsers

		if err := smapping.FillStruct(&data, smapping.MapFields(user)); err != nil {
			log.Error(err.Error())
		} 
		
		users = append(users, data)
	}

	var totalData int64 = 0

	totalData = svc.model.GetTotalDataUsers()

	return users, totalData
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

func (svc *service) ModifyUser(userData dtos.UpdateUser, UserID int) bool {
	newData := user.User{}
	
	err := smapping.FillStruct(&newData, smapping.MapFields(userData))
	if err != nil {
		log.Error(err)
		return false
	}

	newData.ID = UserID
	if userData.Password != ""{
		newData.Password = svc.hash.HashPassword(userData.Password)
	}
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

func (svc *service) MyProfile(UserID int) *dtos.ResMyProfile {
	res := dtos.ResMyProfile{}

	user := svc.model.SelectByID(UserID)
	
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

func (svc *service) SearchUsersByUsername(username string, page, size int) ([]dtos.ResGetAllUsers, int64) {
	usersList := svc.model.FindUsername(username, page, size)

	resUsersList := make([]dtos.ResGetAllUsers, len(usersList))
	for i, n := range usersList {
		resUsersList[i] = dtos.ResGetAllUsers{
			ID: n.ID,
			Username: n.Username,
			RoleID: n.RoleID,
			ExpiryDate: n.ExpiryDate,
		}
	}

	var totalData int64 = 0

	totalData = svc.model.GetTotalDataUsers()
	
	return resUsersList, totalData
}