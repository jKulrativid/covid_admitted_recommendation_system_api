package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (service *UserService) Register(newUser *entities.User) error {
	handleError := service.userRepo.Register(newUser)
	return handleError
}

func (service *UserService) SignIn(newUser *entities.User) (string, error) {
	if err := service.userRepo.Validate(newUser); err != nil {
		return "", err

	}
	return "JWT TOKEN", nil

}
