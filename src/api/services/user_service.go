package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"strconv"
	"time"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}

func (service *UserService) Register(newUser *entities.User) error {
	handleError := service.userRepo.Register(newUser)
	return handleError

}

func (service *UserService) SignIn(newUser *entities.User) error {
	if err := service.verifyUser(newUser); err != nil {
		return err

	}
	return nil

}

func (service *UserService) SignOut(newUser *entities.User) error {
	if err := service.verifyUser(newUser); err != nil {
		return err

	}
	return nil

}

func (service *UserService) CreateAuth(userID uint64, td *TokenDetail) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	stringID := strconv.Itoa(int(userID))

	errAccess := service.userRepo.AddTokenToClient(td.AccessToken, stringID, at.Sub(now))
	if errAccess != nil {
		return errAccess

	}
	errRefresh := service.userRepo.AddTokenToClient(td.RefreshToken, stringID, rt.Sub(now))
	if errRefresh != nil {
		return errRefresh

	}
	return nil

}

func (service *UserService) verifyUser(newUser *entities.User) error {
	// call userRepo to pull userdata from database here
	return nil

}
