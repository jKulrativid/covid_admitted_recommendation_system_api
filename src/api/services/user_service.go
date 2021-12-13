package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
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

func (service *UserService) verifyUser(newUser *entities.User) error {
	// call userRepo to pull userdata from database here
	return nil

}
