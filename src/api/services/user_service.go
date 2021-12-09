package services

import (
	"covid_admission_api/entities"
	"covid_admission_api/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo: *repo,
	}
}

func (service *UserService) Register(newUser *entities.User) error {
	handleError := service.userRepo.Register(newUser)
	return handleError
}
