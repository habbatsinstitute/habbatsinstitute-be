package usecase

import (
	"errors"
	"institute/features/auth"
	"institute/features/auth/dtos"
	"institute/helpers"
	"os"
	"strconv"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/mashingan/smapping"
	"github.com/sirupsen/logrus"
)

type service struct {
	model auth.Repository
	jwt helpers.JWTInterface
	hash helpers.HashInterface
	validator helpers.ValidationInterface
}

func New(model auth.Repository, jwt helpers.JWTInterface, hash helpers.HashInterface, validator helpers.ValidationInterface) auth.Usecase {
	return &service{
		model: model,
		jwt: jwt,
		hash: hash,
		validator: validator,
	}
}

func (svc *service) Register(newData dtos.InputUser) (*dtos.ResUser, []string, error){
	errMap := svc.validator.ValidateRequest(newData)
	if errMap != nil {
		return nil, errMap, nil
	}

	if errorList, err := svc.ValidateInput(newData); err  != nil || len(errorList) > 0 {
		return nil, errorList, err
	}

	newUser := auth.User{}
	err := smapping.FillStruct(&newUser, smapping.MapFields(newData))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	checkUser, _ := svc.model.SelectByUsername(newUser.Username)
	if checkUser != nil {
		logrus.Print("User already exists")
		return nil, nil, errors.New("user already exists")
	}

	newUser.Password = svc.hash.HashPassword(newUser.Password)
	userModel, err := svc.model.Register(&newUser)
	if userModel == nil {
		return nil, nil, err
	}

	resCustomer := dtos.ResUser{}
	err = smapping.FillStruct(&resCustomer, smapping.MapFields(userModel))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	return &resCustomer, nil, nil
}

func (svc *service) Login(data dtos.RequestLogin) (*dtos.LoginResponse, []string, error) {
	errMap := svc.validator.ValidateRequest(data)
	if errMap != nil {
		return nil, errMap, nil
	}

	user, err := svc.model.Login(data.Username)
		if user.RoleID == 1 && user.ExpiryDate.Before(time.Now()){
			return nil, nil, errors.New("account is expired")
	}
	if err != nil {
		return nil, nil, err
	}

	if !svc.hash.CompareHash(data.Password, user.Password) {
		return nil, nil, errors.New("invalid password")
	}

	resUser := dtos.LoginResponse{}

	err = smapping.FillStruct(&resUser, smapping.MapFields(user))
	if err != nil {
		log.Error(err)
		return nil, nil, err
	}

	userID := strconv.Itoa(user.ID)
	roleID := strconv.Itoa(resUser.RoleID)
	tokenData := svc.jwt.GenerateJWT(userID, roleID)

	if tokenData == nil {
		log.Error("Token process failed")
		return nil, nil, errors.New("generate token failed")
	}

	resUser.AccessToken = tokenData["access_token"].(string)
	resUser.RefreshToken = tokenData["refresh_token"].(string)

	return &resUser, nil, nil
}

func (svc *service) RefreshJWT(jwt dtos.RefreshJWT) (*dtos.ResJWT, error){
	parsedRefreshToken, err := svc.jwt.ValidateToken(jwt.RefreshToken, os.Getenv("REFSECRET"))
	if err != nil {
		return nil, errors.New("validate token failed")
	}

	token := svc.jwt.RefereshJWT(parsedRefreshToken)
	if token == nil {
		return nil, errors.New("refresh jwt failed")
	}

	var resJWT  dtos.ResJWT
	resJWT.AccessToken = token["access_token"].(string)
	resJWT.RefreshToken = token["refresh_token"].(string)
	
	return &resJWT, nil
}

func (svc *service) ValidateInput( input dtos.InputUser) ([]string, error){
	var errorList []string

	if errMap := svc.validator.ValidateRequest(input); errMap != nil {
		errorList = append(errorList, errMap...)
	}
	if len(input.Username) <= 5 {
		errorList = append(errorList, "username must be at least 5 characters")
	}
	if input.ExpiryDate.Before(time.Now()){
		errorList = append(errorList, "expiry date must be greater than today")
	}
	return errorList, nil
}