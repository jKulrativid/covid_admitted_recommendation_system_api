package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
)

type UserService struct {
	userRepo     repositories.UserRepository
	jwtTokenRepo repositories.JWTTokenRepository
}

func NewUserService(u repositories.UserRepository, jwt repositories.JWTTokenRepository) *UserService {
	return &UserService{
		userRepo:     u,
		jwtTokenRepo: jwt,
	}
}

func (service *UserService) Register(newUser *entities.User) error {
	handleError := service.userRepo.Register(newUser)
	return handleError

}

func (service *UserService) SignIn(newUser *entities.User) (string, error) {
	if err := service.verifyUser(newUser); err != nil {
		return "", err

	}
	token := "JWT_TOKEN"
	return token, nil

}

func (service *UserService) SignOut(newUser *entities.User) error {
	if err := service.verifyUser(newUser); err != nil {
		return err

	}
	// deactivate user token here
	return nil

}

func (service *UserService) verifyUser(newUser *entities.User) error {
	// call userRepo to pull userdata from database here
	return nil

}
