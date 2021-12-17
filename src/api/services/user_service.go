package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
	"fmt"
	"time"

	"github.com/twinj/uuid"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(u repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: u,
	}
}

func (service *UserService) Register(newUser *entities.UserRegister) (handleError error) {
	user := &entities.User{
		Uuid:           uuid.NewV4().String(),
		UserName:       newUser.UserName,
		Email:          newUser.Email,
		HashedPassword: newUser.Password,
		Salt:           "123456lol",
	}
	handleError = service.userRepo.Register(user)
	return handleError

}

func (service *UserService) SignIn(user *entities.UserSignIn) (uuid string, err error) {
	if uuid, err = service.verifyUser(user); err != nil {
		return "", err

	}
	return uuid, nil

}

func (service *UserService) SignOut(u *entities.User) (err error) {
	err = nil
	return err

}

func (service *UserService) CreateAuth(uuid string, td *TokenDetail) (err error) {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	stringID := uuid

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

func (service *UserService) FetchAuth(authD *AccessDetails) (string, error) {
	uuid, err := service.userRepo.GetFromClient(authD.AccessUuid)
	if err != nil {
		return "", err
	}
	return uuid, nil
}

func (service *UserService) verifyUser(userSignIn *entities.UserSignIn) (uuid string, err error) {
	// call userRepo to pull userdata from database here
	var user entities.User
	err = service.userRepo.PullUserData(&user, userSignIn.UserName)
	if err != nil {
		return "", fmt.Errorf("username not found")
	}
	// TODO hashed before check password
	if user.HashedPassword != userSignIn.Password {
		return "", fmt.Errorf("invalid username or password")
	}
	return user.Uuid, nil
}
