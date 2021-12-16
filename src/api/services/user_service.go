package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"fmt"
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

func (service *UserService) Register(newUser *entities.UserRegister) error {
	handleError := service.userRepo.Register(newUser)
	return handleError

}

func (service *UserService) SignIn(user *entities.UserSignIn) error {
	if err := service.verifyUser(user); err != nil {
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

func (service *UserService) CreateAuth(userID uint64, td *TokenDetail) (err error) {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	stringID := strconv.Itoa(int(userID))

	err = service.userRepo.AddTokenToClient(td.AccessToken, stringID, at.Sub(now))
	if err != nil {
		return err

	}
	err = service.userRepo.AddTokenToClient(td.RefreshToken, stringID, rt.Sub(now))
	if err != nil {
		return err

	}
	return nil

}

func (service *UserService) verifyUser(userSignIn *entities.UserSignIn) (err error) {
	// call userRepo to pull userdata from database here
	var user entities.User
	err = service.userRepo.PullUserData(&user, userSignIn.UserName)
	if err != nil {
		return err
	}
	// TODO hashed before check password
	if user.UserName != userSignIn.UserName || user.HashedPassword != userSignIn.Password {
		return fmt.Errorf("invalid username or password")
	}
	return nil
}
